function Footer() {
  return (
    <footer className="bg-blue-500">
      <div className="w-full mx-auto max-w-screen p-3 flex items-center justify-center">
        <span className="text-sm text-white text-center">
          © 2024<span className="text-sm text-white"> </span>
          <a
            href="https://biocorellc.com/"
            target="_blank"
            className="hover:underline"
          >
            Biomechanics Consulting & Research, LLC
          </a>
          . All Rights Reserved.
        </span>
      </div>
    </footer>
  );
}

export default Footer;
