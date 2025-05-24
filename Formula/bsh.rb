class Bsh < Formula
  desc "Blockchain Server Hosting CLI"
  homepage "https://github.com/kimtuna/bsh"
  version "0.1.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/kimtuna/bsh/releases/download/v0.1.0/bsh-darwin-arm64.tar.gz"
      sha256 "..." # 실제 sha256 값으로 교체 필요
    else
      url "https://github.com/kimtuna/bsh/releases/download/v0.1.0/bsh-darwin-amd64.tar.gz"
      sha256 "..." # 실제 sha256 값으로 교체 필요
    end
  end

  def install
    bin.install "bsh"
  end

  test do
    system "#{bin}/bsh", "--version"
  end
end 