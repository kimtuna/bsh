class Bsh < Formula
  desc "Blockchain Server Hosting CLI"
  homepage "https://github.com/kimtuna/bsh"
  version "0.1.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/kimtuna/bsh/releases/download/v0.1.0/bsh-darwin-arm64.tar.gz"
      sha256 "f8771fbe5d8058641f898d4b98dc91b707dfafe03ab9136dd40dfe6a7c2bacbb"
   
    end
  end

  def install
    bin.install "bsh"
  end

  test do
    system "#{bin}/bsh", "--version"
  end
end 