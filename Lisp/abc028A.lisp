(defun main ()
  (setf a (read))
  (when (<= a 59) (princ "Bad"))
  (when (and (<= 60 a) (<= a 89)) (princ "Good"))
  (when (and (<= 90 a) (<= a 99)) (princ "Great"))
  (when (= a 100) (princ "Perfect"))
  (fresh-line))
(main)