(defun main ()
  (let ((n (read)) (i 1))
  (write n)
  (fresh-line)
  (loop
    (if (> i n)(return 0))
    (setq i (+ i 1))
    (write 1)
    (fresh-line))))
(main)