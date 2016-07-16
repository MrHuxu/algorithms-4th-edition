object Exercise7 {
  class Stack[T](var data: List[T]) {
    def push(arg: T): Stack[T] = {
      this.data = arg +: this.data
      return this
    }

    def pop(): String = {
      val top = this.data match {
        case head +: tail => {
          this.data = tail
          s"$head"
        }
        case Nil => "Nil"
      }
      return top
    }

    def peek(): String = {
      val top = this.data match {
        case head +: tail => s"$head"
        case Nil => "Nil"
      }
      return top
    }
  }

  def main(args: Array[String]): Unit = {
    var a = new Stack[String](List())
    a.push("it").push("was").push("the").push("best").push("of").push("times")
    println(a.peek)
    println(a.peek)
    println(a.pop)
    println(a.peek)
  }
}
