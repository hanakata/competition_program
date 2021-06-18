(defun main ()
  (let ((n (read)) (i 0))
  (let ((str (make-array 5 :initial-contents '(A B C D E))))
  (loop
    (if (string= n (aref str i))(return i))
    (setq i (+ i 1)))
  (write (+ i 1))))
  (fresh-line))
(main)
