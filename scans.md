        Для чтения входных данных необходимо получить их
        из стандартного потока ввода.
        Данные во входном потоке соответствуют описанному
        в условии формату. Обычно входные данные состоят
        из нескольких строк.
        Можно использовать несколько методов:
        * fmt.Scanln(), fmt.Scan(), fmt.Scanf() -- функции, позволяющие считывать отдельные слова в отдельные переменные;
        * bufio.Reader -- позволяет читать строку входного текста или один символ;
        * bufio.Scanner -- предоставляет удобный интерфейс для чтения строк текста.
        Чтобы прочитать из строки стандартного потока:
        * число --
        var number int64
        _, err := fmt.Scan(&number)
        if err != nil {
            log.Fatal(err)
        }
        * строку --
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        err := scanner.Err()
        if err != nil {
            log.Fatal(err)
        }
        * массив чисел известной длины --
        var arr = make([]int, 10)
        for i := 0; i < len(arr); i++ {
            _, err := fmt.Scan(&arr[i])
            if err != nil {
                log.Fatal(err)
            }
        }
        * последовательность слов до конца файла --
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Split(bufio.ScanWords)
        var words []string
        for scanner.Scan() {
            words = append(words, scanner.Text())
        }
        Чтобы вывести результат в стандартный поток вывода,
        можно использовать fmt.Println() или fmt.Print().
        Возможное решение задачи "Вычислите сумму A+B":
        var a, b int64
        _, err := fmt.Scan(&a)
        if err != nil {
            log.Fatal(err)
        }
        _, err = fmt.Scan(&b)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(a + b)
