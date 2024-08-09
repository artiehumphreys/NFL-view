import bpy  # type: ignore
import os
import mathutils  # type: ignore
from pathlib import Path
import sys
import csv


class bvh_to_glb:

    def __init__(
        self,
        dir: str,
        output_path: str,
        field_obj: str = "rendering/court.obj",
    ) -> None:
        self.output_path: str = output_path
        self.filename: str = ""
        self.player_ids: dict[int, str] = {}
        self.teams: dict[str, int] = {}
        self.scale: float = 1.25
        self.dir: str = dir
        self.field_obj = field_obj
        self.start_frame, self.end_frame = 0, 0
        self.process_players(self.dir)
        self.count = 0

    def get_team(self) -> str:
        """Getter method for fetching player team."""
        split: list[str] = self.filename.split("_")
        return split[1]

    def get_player_name(self) -> str:
        """Getter method for fetching player name based on unique id. This feature is not required."""
        try:
            split: list[str] = self.filename.split("_")
            return self.player_ids[int(split[2])]
        except:
            return ""

    def process_players(self, directory: str) -> None:
        """Initalize dictionary for player name mappings based on file name."""
        for file in os.listdir(directory):
            if not file.endswith(".bvh"):
                continue
            filename: str = Path(file).stem
            split: list[str] = filename.split("_")
            if not split[1] in self.teams:
                self.teams[split[1]] = len(self.teams)
            self.player_ids[int(split[2])] = ""

    def read_ball_csv(self, file="data/Ball_Track.csv"):
        frame = 1
        with open(file, mode="r") as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                self.ball_track[frame] = (
                    float(row["X"]),
                    float(row["Y"]),
                    float(row["Z"]),
                )
                frame += 1

    def assign_team_color(self, shape: bpy.types.Object) -> None:
        """Based on number assigned to team of which a player is on (either a 0 or 1), color the bone."""
        team: str = self.get_team()
        color: float = self.teams[team]
        mat: bpy.types.Material = bpy.data.materials.new(name=f"Material_{team}")
        mat.diffuse_color = (color, color, 1, 1)
        shape.data.materials.append(mat)

    def create_sphere_at_bone(
        self, bone: bpy.types.Bone, armature: bpy.types.Object
    ) -> None:
        """Create the bone object. If the bone is a hand, elongate it to resemble humans."""
        if bone.name == "lWrist" or bone.name == "rWrist":
            bpy.ops.mesh.primitive_uv_sphere_add(
                radius=bone.head_radius / self.scale, location=bone.head_local
            )
            sphere: bpy.types.Object = bpy.context.object
            sphere.scale = ((bone.head_local - bone.tail_local).length * 10, 1, 1)
        else:
            bpy.ops.mesh.primitive_uv_sphere_add(
                radius=bone.head_radius, location=bone.head_local
            )
        sphere = bpy.context.object
        sphere.name = f"sphere_{bone.name}_{self.get_player_name()}"
        if bone.name == "baseHead":
            head_height = 0.15
            sphere.location += mathutils.Vector((0, 0, head_height))
            sphere.scale = (self.scale, self.scale, self.scale)

        mod: bpy.types.Modifier = sphere.modifiers.new(name="Armature", type="ARMATURE")
        mod.object = armature
        sphere.parent = armature
        sphere.parent_type = "BONE"
        sphere.parent_bone = bone.name
        self.assign_team_color(sphere)

    def create_bone_connector(
        self, bone: bpy.types.Bone, armature: bpy.types.Object
    ) -> None:
        """Create the bone connectors (arms, legs, etc.) for eahc of the bones."""
        parent: bpy.types.Bone = bone.parent

        cone_vector: mathutils.Vector = bone.head_local - parent.head_local
        cone_length: float = cone_vector.length
        cone_direction: mathutils.Vector = cone_vector.normalized()

        cone_midpoint_world: mathutils.Vector = (
            parent.head_local + bone.head_local
        ) / 2

        bpy.ops.mesh.primitive_cone_add(
            radius1=bone.head_radius / self.scale,
            radius2=bone.tail_radius / self.scale,
            depth=cone_length,
            location=cone_midpoint_world,
        )
        cone: bpy.types.Object = bpy.context.object
        cone.name = f"Cone_{parent.name}_to_{bone.name}_{self.filename}"

        rotation: mathutils.Euler = cone_direction.to_track_quat("Z", "Y").to_euler()
        cone.rotation_euler = rotation

        cone.parent = armature
        cone.parent_type = "BONE"
        cone.parent_bone = parent.name
        self.assign_team_color(cone)

    def display_ball(self):
        bpy.ops.object.select_all(action="DESELECT")
        bpy.ops.mesh.primitive_uv_sphere_add(radius=0.124, location=(0, 0, 0))
        ball_obj = bpy.context.active_object
        ball_obj.name = "ball"
        ball_obj.data.materials.clear()
        material = bpy.data.materials.new(name="ball color")
        material.diffuse_color = (1, 0.3, 0.1, 1)
        ball_obj.data.materials.append(material)

        for frame in range(0, len(self.ball_track)):
            bpy.context.scene.frame_set(frame)
            ball_obj.location = self.ball_track[frame]
            ball_obj.keyframe_insert(data_path="location", index=-1)

    def convert_bvh_to_glb(self, output_name: str) -> None:
        """Driver method for conversion."""
        # display_ball = False
        bpy.ops.wm.read_factory_settings(use_empty=True)

        try:
            bpy.ops.wm.obj_import(filepath=self.field_obj)
        except:
            print(
                "If you would like to include a field/court in the animation, ensure that the path is correct. 'rendering/court.obj'"
            )
        if not os.path.exists(self.dir):
            print(
                "Please ensure that the directory that contains all of the .bvh files is correctly cited."
            )
            exit(1)
        for file in os.listdir(self.dir):
            if not file.endswith(".bvh"):
                print(f"Ignoring file: {file}. It is not a .bvh file.")
                continue
            self.filename = Path(file).stem

            bvh_path = os.path.join(self.dir, file)
            bpy.ops.import_anim.bvh(filepath=bvh_path)
            bpy.context.scene.render.fps = 20

            armature: bpy.types.Object = bpy.context.object
            if armature.type != "ARMATURE":
                continue
            bpy.context.view_layer.objects.active = armature

            action: bpy.types.Action = armature.animation_data.action
            if not action:
                print("Error getting frame range. Try manually inputting it")
                exit(1)
            self.start_frame = int(action.frame_range[0])
            self.end_frame = int(action.frame_range[1])

            # if not display_ball:
            #     self.display_ball()
            #     display_ball = True

            for bone in armature.pose.bones:
                bone.bone.use_local_location = True
                bone.bone.use_relative_parent = True
                self.create_sphere_at_bone(bone.bone, armature)
                if bone.bone.parent:
                    self.create_bone_connector(bone.bone, armature)

        bpy.ops.object.select_all(action="SELECT")

        bpy.ops.export_scene.gltf(
            filepath=f"{self.output_path}/{sys.argv[1]}_{sys.argv[2]}",
            export_format="GLB",
        )


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Not enough arguments provided.")
        sys.exit(2)

    game_id = sys.argv[1]
    play_id = sys.argv[2]
    pad = play_id.rjust(6, "0")

    bvh_dir = os.path.abspath(
        f"../public/alpha/reconstruction/{game_id}/{pad}/final_results/bvh"
    )
    output_path = os.path.abspath("../public/alpha/final_files")

    if not os.path.exists(f"{output_path}/{game_id}_{play_id}.glb"):
        field_obj = (
            sys.argv[3]
            if len(sys.argv) > 3
            else os.path.abspath("../public/alpha/input_rendering/NFLstadium.obj")
        )

        if not os.path.exists(bvh_dir):
            print(f"BVH directory does not exist: {bvh_dir}")
            sys.exit(2)

        if not os.path.exists(output_path):
            os.mkdir(output_path)

        converter = bvh_to_glb(
            dir=bvh_dir,
            output_path=output_path,
            field_obj=field_obj,
        )

        converter.convert_bvh_to_glb(converter.output_path + "/Pose3D")
