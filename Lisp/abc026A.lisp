(defun main ()
  (setf a (read))
  (write (* (/ a 2) (/ a 2)))
  (fresh-line))
(main)