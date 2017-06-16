(define match
  (lambda (regex text)
    (let ([regex (string->list regex)]
	  [text (string->list text)])
      (matchregex regex text))))


(define matchregex
  (lambda (regex text)
    (cond
     [(null? regex) #f]
     [(eq? (car regex) #\^)
      (matchhere (cdr regex) text)]
     [(matchhere regex text) #t]
     [else
      (let loop ([text text])
	(cond
	 [(null? text) #f]
	 [else
	  (if (matchhere regex text)
	    #t
	    (loop (cdr text)))]))])))


(define matchhere
  (lambda (regex text)
    (cond
     [(null? regex) #t]
     [(and (eq? (car regex) #\$) (null? (cdr regex)))
      (null? text)]
     [(null? text) #f]
     [(and (not (null? (cdr regex))) (eq? (cadr regex) #\*))
      (matchstar (car regex) (cddr regex) text)]
     [(and (not (null? (cdr regex))) (eq? (cadr regex) #\+))
      (matchplus (car regex) (cddr regex) text)]
     [(or (eq? (car regex) #\.) (eq? (car regex) (car text)))
      (matchhere (cdr regex) (cdr text))]
     [else #f])))


(define matchstar
  (lambda (c regex text)
    (cond
     [(matchhere regex text) #t]
     [else
      (let loop ([text text])
	(cond
	 [(or (null? text) (and (not (eq? (car text) c)) (not (eq? c #\.)))) #f]
	 [else
	  (if (matchhere regex text)
	      #t
	      (loop (cdr text)))]))])))


(define matchplus
  (lambda (c regex text)
    (let loop ([text text])
      (cond
       [(or (null? text) (and (not (eq? (car text) c)) (not (eq? c #\.)))) #f]
       [else
	(if (matchhere regex text)
	    #t
	    (loop (cdr text)))]))))
