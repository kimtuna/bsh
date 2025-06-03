class Bsh < Formula
  desc "Blockchain Server Hosting CLI"
  homepage "https://github.com/kimtuna/bsh"
  version "0.1.3"
  url "https://github.com/kimtuna/bsh/releases/download/v0.1.3/bsh-darwin-arm64.tar.gz"
  sha256 "84299584d26e2357b4e3c35c695c44cd182fdcab7a48541a63bd5a488413d167"
  license "MIT"

  def install
    bin.install "bsh"
  end

  test do
    system "#{bin}/bsh", "--version"
  end
end 