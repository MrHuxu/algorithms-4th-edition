object Exercise19 {
  sealed trait LinkedList[+E]

  case class Node[+E]( val head : E, val tail : LinkedList[E]  ) extends LinkedList[E]

  case object Empty extends LinkedList[Nothing]

  def main(args: Array[String]): Unit = {
  }
}
