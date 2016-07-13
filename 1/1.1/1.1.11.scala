object Exercise11 {
  def main(args: Array[String]) = {
    var arr = Array(Array(true, false, true), Array(true, true, false), Array(false, true, true))
    println(" 123")
    for (i <- 0 until arr.length) {
      var row: String = (i + 1).toString
      arr(i).foreach(elem => row = row + (if (elem) "x" else "_"))
      println(row)
    }
  }
}