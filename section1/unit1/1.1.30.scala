object Exercise30 {
  def gcd(a: Int, b: Int): Int = {
    if (b == 0) return a
    return gcd(b, a % b)
  }

  def main(args: Array[String]) = {
    println(gcd(35, 25))
    println(gcd(36, 28))
    println(args(0), args(1))
  }
}