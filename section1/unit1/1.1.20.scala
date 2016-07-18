import scala.math._

object Exercise20 {
  def calc(res: Integer, num: Integer): Unit = {
    if (num == 1) {
      println(log(res.toDouble))
    } else {
      calc(res * num, num - 1)
    }
  }

  def main(args: Array[String]) = {
    calc(1, 5)
  }
}