class Bsh < Formula
  desc "Blockchain Server Hosting CLI"
  homepage "https://github.com/kimtuna/bsh"
  version "0.1.0"
  
  if OS.mac?
    url "https://github.com/kimtuna/bsh/releases/download/v#{version}/bsh-darwin-amd64.tar.gz"
    sha256 "b50f4d04e0a20ff4a0c55ecbb90f3cb852bb6c702c3d8cd95ef1747465a099e5"
  elsif OS.linux?
    url "https://github.com/kimtuna/bsh/releases/download/v#{version}/bsh-linux-amd64.tar.gz"
    sha256 "4465ee71398660a50d4ef264e884b10367fe1ca0f03d4ca8454fa93d8452bc49"
  end

  def install
    bin.install "bsh"
  end

  test do
    system "#{bin}/bsh", "--version"
  end
end 