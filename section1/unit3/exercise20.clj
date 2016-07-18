(ns exercise19
  (load-file "list-node.clj"))

(import unit3.list_node.ListNode)

(def f (ListNode. 6 nil))
(def e (ListNode. 5 f))
(def d (ListNode. 4 e))
(def c (ListNode. 3 d))
(def b (ListNode. 2 c))
(def a (ListNode. 1 b))

(defn delete
  [head k]
  )
