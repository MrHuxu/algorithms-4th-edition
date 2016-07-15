object Excecise16 {
  class Rational(var numerator: Int, var denominator: Int) {
    def gcd(a: Int, b: Int): Int = {
      if (b == 0) return a
      return gcd(b, a % b)
    }

    def plus(b: Rational): Rational = {
      var newNumerator: Int = this.numerator * b.denominator + b.numerator * this.denominator
      var newDenominator: Int = this.denominator * b.denominator
      var newGcd = this.gcd(newNumerator, newDenominator)
      return new Rational(newNumerator / newGcd, newDenominator / newGcd)
    }

    def minus(b: Rational): Rational = {
      var newNumerator: Int = this.numerator * b.denominator - b.numerator * this.denominator
      var newDenominator: Int = this.denominator * b.denominator
      var newGcd = this.gcd(newNumerator, newDenominator)
      return new Rational(newNumerator / newGcd, newDenominator / newGcd)
    }

    def times(b: Rational): Rational = {
      var newNumerator: Int = this.numerator * b.numerator
      var newDenominator: Int = this.denominator * b.denominator
      var newGcd = this.gcd(newNumerator, newDenominator)
      return new Rational(newNumerator / newGcd, newDenominator / newGcd)
    }

    def divides(b: Rational): Rational = {
      var newNumerator: Int = this.numerator * b.denominator
      var newDenominator: Int = this.denominator * b.numerator
      var newGcd = this.gcd(newNumerator, newDenominator)
      return new Rational(newNumerator / newGcd, newDenominator / newGcd)
    }

    def equals(b: Rational): Boolean = {
      var gcdA = this.gcd(this.numerator, this.denominator)
      var gcdB = this.gcd(b.numerator, b.denominator)
      return this.numerator / gcdA == b.numerator / gcdB && this.denominator / gcdA == b.denominator / gcdB
    }

    override def toString(): String = {
      return this.numerator.toString + "/" + this.denominator.toString
    }
  }

  def main(args: Array[String]): Unit = {
    var a = new Rational(2, 3)
    var b = new Rational(5, 6)
    println((a plus b).toString)
    println((a minus b).toString)
    println((a times b).toString)
    println((a divides b).toString)
    println(a equals b)
    println(a equals (new Rational(4, 6)))
  }
}
