(defun main ()
  (setf a (concatenate 'list(read-line)))
  (setf b (read))
  (setf p (floor(/ (- b 1) 5.0)))
  (setf q (mod b 5))
  (if (= q 0) (format t"~A~A" (nth p a)(nth 4 a)) (format t"~A~A" (nth p a)(nth (- q 1) a)))
  (fresh-line))
(main)
