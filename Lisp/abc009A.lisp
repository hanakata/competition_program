(defun main ()
  (let ((a (read)))
  (write (+(floor (/ a 2))  (mod a 2))))
  (fresh-line))
(main)