(defprotocol PStack
  "A stack protocal"
  (push [this val] "Push element into the stack")
  (pop [this] "Pop element from stack")
  (top [this] "Get top element from stack"))

(defrecord Stack [coll]
  PStack
  (push [_ val]
    (swap! coll conj val))
  (pop [_]
    (let [ret (first @coll)]
      (swap! coll rest)
    ret))
  (top [_]
    (first @coll)))

(def s (Stack. (atom '())))

(println (push s 10))

(println (push s 20))

(println (top s))

(println s)

(println (pop s))
