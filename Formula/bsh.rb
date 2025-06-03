class Bsh < Formula
  desc "Blockchain Server Hosting CLI"
  homepage "https://github.com/kimtuna/bsh"
  version "0.1.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/kimtuna/bsh/releases/download/v0.1.1/bsh-darwin-arm64.tar.gz"
      sha256 "ebd8c379fc3241585f18ed343854078d6cf27c324b7f4dc3b99c9d4d8617518b"
   
    end
  end

  def install
    bin.install "bsh"
  end

  test do
    system "#{bin}/bsh", "--version"
  end
end 