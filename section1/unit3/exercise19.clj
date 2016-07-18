(ns exercise19
  (load-file "list-node.clj"))

(import unit3.list_node.ListNode)

(def c (ListNode. 3 nil))
(def b (ListNode. 2 c))
(def a (ListNode. 1 b))

(println (:val a))
(println (-> a :next :next :val))

(def level [])
(def head a)
(while (boolean (:next a)) (do
  (def level (conj level :next))
  (def a (:next a))))

(try
  (println level)
  (def head (assoc-in head level nil))
  (println head)
(catch Exception e
  (println "caught exception" (.getMessage e))))
