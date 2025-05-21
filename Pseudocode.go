program main

kamus
    TipeData Mahasiswa :
        NIM    : string
        Name   : string
        Course : string
        Origin : string

    TipeData Buku :
        Author              : string
        Judul               : string
        Tahun               : integer
        TanggalPeminjaman   : string
        TanggalPengembalian : string

    TipeData Pegawai :
        NIP     : string
        Nama    : string
        Jabatan : string

    studentData   : array of Mahasiswa
    bookData      : array of Buku
    employeeData  : array of Pegawai

algoritma:
    while true do
        tampilkan menu utama
        input(choice)
        switch choice
            case "1":
                addData()
            case "2":
                viewData()
            case "3":
                updateData()
            case "4":
                deleteData()
            case "5":
                searchData()
            case "6":
                tampilkan "Exiting program."
                exit
            default:
                tampilkan "Invalid choice."
        endswitch
    endwhile
endprogram

procedure addData()
algoritma:
    while true do
        tampilkan menu tambah data
        input(choice)
        switch choice
            case "1":
                deklarasi m : Mahasiswa
                tampilkan "ADD MAHASISWA DATA"
                repeat
                    input(m.NIM)
                    if not isNumeric(m.NIM) then
                        tampilkan "Error: NIM must be numeric."
                    endif
                until isNumeric(m.NIM)
                input(m.Name)
                input(m.Course)
                input(m.Origin)
                tambah m ke studentData
                tampilkan "Mahasiswa data successfully added."
            case "2":
                deklarasi b : Buku
                tampilkan "ADD BUKU DATA"
                input(b.Author)
                input(b.Judul)
                input(b.Tahun)
                input(b.TanggalPeminjaman)
                input(b.TanggalPengembalian)
                tambah b ke bookData
                tampilkan "Buku data successfully added."
            case "3":
                deklarasi p : Pegawai
                tampilkan "ADD PEGAWAI DATA"
                input(p.NIP)
                input(p.Nama)
                input(p.Jabatan)
                tambah p ke employeeData
                tampilkan "Pegawai data successfully added."
            case "4":
                return
            default:
                tampilkan "Invalid choice."
        endswitch
    endwhile
endprocedure

function isNumeric(input : string) : boolean
algoritma:
    return true jika input hanya terdiri dari digit
endfunction

procedure viewData()
algoritma:
    while true do
        tampilkan menu view data
        input(choice)
        switch choice
            case "1":
                if studentData kosong then
                    tampilkan "No Mahasiswa data available."
                    continue
                endif
                while true do
                    tampilkan menu sort mahasiswa
                    input(sortChoice)
                    if sortChoice == "3" then
                        break
                    else if sortChoice == "1" then
                        selectionSortMahasiswa(studentData, true)
                    else if sortChoice == "2" then
                        insertionSortMahasiswa(studentData, false)
                    else
                        tampilkan "Invalid sort choice."
                        continue
                    endif
                    tampilkan header tabel mahasiswa
                    for i = 0 to length(studentData)-1 do
                        tampilkan data mahasiswa ke-i
                    endfor
                    tampilkan footer tabel mahasiswa
                    break
                endwhile
            case "2":
                if bookData kosong then
                    tampilkan "No Buku data available."
                    continue
                endif
                while true do
                    tampilkan menu sort buku
                    input(sortChoice)
                    if sortChoice == "3" then
                        break
                    else if sortChoice == "1" then
                        selectionSortBuku(bookData, true)
                    else if sortChoice == "2" then
                        insertionSortBuku(bookData, false)
                    else
                        tampilkan "Invalid sort choice."
                        continue
                    endif
                    tampilkan header tabel buku
                    for i = 0 to length(bookData)-1 do
                        tampilkan data buku ke-i
                    endfor
                    tampilkan footer tabel buku
                    break
                endwhile
            case "3":
                if employeeData kosong then
                    tampilkan "No Pegawai data available."
                    continue
                endif
                while true do
                    tampilkan menu sort pegawai
                    input(sortChoice)
                    if sortChoice == "3" then
                        break
                    else if sortChoice == "1" then
                        selectionSortPegawai(employeeData, true)
                    else if sortChoice == "2" then
                        insertionSortPegawai(employeeData, false)
                    else
                        tampilkan "Invalid sort choice."
                        continue
                    endif
                    tampilkan header tabel pegawai
                    for i = 0 to length(employeeData)-1 do
                        tampilkan data pegawai ke-i
                    endfor
                    tampilkan footer tabel pegawai
                    break
                endwhile
            case "4":
                return
            default:
                tampilkan "Invalid choice."
        endswitch
    endwhile
endprocedure

procedure selectionSortMahasiswa(data : array of Mahasiswa, ascending : boolean)
algoritma:
    for i = 0 to length(data)-2 do
        idx <- i
        for j = i+1 to length(data)-1 do
            if ascending and data[j].NIM < data[idx].NIM then
                idx <- j
            endif
        endfor
        tukar data[i] dengan data[idx]
    endfor
endprocedure

procedure insertionSortMahasiswa(data : array of Mahasiswa, descending : boolean)
algoritma:
    for i = 1 to length(data)-1 do
        key <- data[i]
        j <- i-1
        while j >= 0 and descending and data[j].NIM < key.NIM do
            data[j+1] <- data[j]
            j <- j-1
        endwhile
        data[j+1] <- key
    endfor
endprocedure

procedure selectionSortBuku(data : array of Buku, ascending : boolean)
algoritma:
    for i = 0 to length(data)-2 do
        idx <- i
        for j = i+1 to length(data)-1 do
            if ascending and data[j].Judul < data[idx].Judul then
                idx <- j
            endif
        endfor
        tukar data[i] dengan data[idx]
    endfor
endprocedure

procedure insertionSortBuku(data : array of Buku, descending : boolean)
algoritma:
    for i = 1 to length(data)-1 do
        key <- data[i]
        j <- i-1
        while j >= 0 and descending and data[j].Judul < key.Judul do
            data[j+1] <- data[j]
            j <- j-1
        endwhile
        data[j+1] <- key
    endfor
endprocedure

procedure selectionSortPegawai(data : array of Pegawai, ascending : boolean)
algoritma:
    for i = 0 to length(data)-2 do
        idx <- i
        for j = i+1 to length(data)-1 do
            if ascending and data[j].NIP < data[idx].NIP then
                idx <- j
            endif
        endfor
        tukar data[i] dengan data[idx]
    endfor
endprocedure

procedure insertionSortPegawai(data : array of Pegawai, descending : boolean)
algoritma:
    for i = 1 to length(data)-1 do
        key <- data[i]
        j <- i-1
        while j >= 0 and descending and data[j].NIP < key.NIP do
            data[j+1] <- data[j]
            j <- j-1
        endwhile
        data[j+1] <- key
    endfor
endprocedure

procedure updateData()
algoritma:
    while true do
        tampilkan menu update data
        input(choice)
        switch choice
            case "1":
                tampilkan "Enter NIM of the Mahasiswa to update:"
                input(nim)
                for i = 0 to length(studentData)-1 do
                    if studentData[i].NIM == nim then
                        tampilkan data mahasiswa ke-i
                        tampilkan menu field update mahasiswa
                        input(field)
                        switch field
                            case "1":
                                input(studentData[i].Name)
                            case "2":
                                input(studentData[i].NIM)
                            case "3":
                                input(studentData[i].Course)
                            case "4":
                                input(studentData[i].Origin)
                            default:
                                tampilkan "Invalid choice."
                        endswitch
                        tampilkan "Mahasiswa data successfully updated."
                        return
                    endif
                endfor
                tampilkan "Mahasiswa with the specified NIM not found."
            case "2":
                tampilkan "Enter Judul of the Buku to update:"
                input(judul)
                for i = 0 to length(bookData)-1 do
                    if lower(bookData[i].Judul) == lower(judul) then
                        tampilkan data buku ke-i
                        tampilkan menu field update buku
                        input(field)
                        switch field
                            case "1":
                                input(bookData[i].Author)
                            case "2":
                                input(bookData[i].Judul)
                            case "3":
                                input(bookData[i].Tahun)
                            case "4":
                                input(bookData[i].TanggalPeminjaman)
                            case "5":
                                input(bookData[i].TanggalPengembalian)
                            default:
                                tampilkan "Invalid choice."
                        endswitch
                        tampilkan "Buku data successfully updated."
                        return
                    endif
                endfor
                tampilkan "Buku with the specified Judul not found."
            case "3":
                tampilkan "Enter NIP of the Pegawai to update:"
                input(nip)
                for i = 0 to length(employeeData)-1 do
                    if employeeData[i].NIP == nip then
                        tampilkan data pegawai ke-i
                        tampilkan menu field update pegawai
                        input(field)
                        switch field
                            case "1":
                                input(employeeData[i].Nama)
                            case "2":
                                input(employeeData[i].NIP)
                            case "3":
                                input(employeeData[i].Jabatan)
                            default:
                                tampilkan "Invalid choice."
                        endswitch
                        tampilkan "Pegawai data successfully updated."
                        return
                    endif
                endfor
                tampilkan "Pegawai with the specified NIP not found."
            case "4":
                return
            default:
                tampilkan "Invalid choice."
        endswitch
    endwhile
endprocedure

procedure deleteData()
algoritma:
    while true do
        tampilkan menu delete data
        input(choice)
        switch choice
            case "1":
                tampilkan "Enter NIM of the Mahasiswa to delete:"
                input(nim)
                for i = 0 to length(studentData)-1 do
                    if studentData[i].NIM == nim then
                        tampilkan data mahasiswa ke-i
                        tampilkan "Delete this data? (yes/no):"
                        input(confirm)
                        if lower(confirm) == "yes" then
                            hapus studentData[i]
                            tampilkan "Mahasiswa data successfully deleted."
                        else
                            tampilkan "Deletion canceled."
                        endif
                        return
                    endif
                endfor
                tampilkan "Mahasiswa with the specified NIM not found."
            case "2":
                tampilkan "Enter Judul of the Buku to delete:"
                input(judul)
                for i = 0 to length(bookData)-1 do
                    if lower(bookData[i].Judul) == lower(judul) then
                        tampilkan data buku ke-i
                        tampilkan "Delete this data? (yes/no):"
                        input(confirm)
                        if lower(confirm) == "yes" then
                            hapus bookData[i]
                            tampilkan "Book data successfully deleted."
                        else
                            tampilkan "Deletion canceled."
                        endif
                        return
                    endif
                endfor
                tampilkan "Buku with the specified Judul not found."
            case "3":
                tampilkan "Enter NIP of the Pegawai to delete:"
                input(nip)
                for i = 0 to length(employeeData)-1 do
                    if employeeData[i].NIP == nip then
                        tampilkan data pegawai ke-i
                        tampilkan "Delete this data? (yes/no):"
                        input(confirm)
                        if lower(confirm) == "yes" then
                            hapus employeeData[i]
                            tampilkan "Pegawai data successfully deleted."
                        else
                            tampilkan "Deletion canceled."
                        endif
                        return
                    endif
                endfor
                tampilkan "Pegawai with the specified NIP not found."
            case "4":
                return
            default:
                tampilkan "Invalid choice."
        endswitch
    endwhile
endprocedure

procedure searchData()
algoritma:
    while true do
        tampilkan menu search data
        input(bukuChoice)
        switch bukuChoice
            case "1":
                availableBooks <- filter bookData dimana TanggalPeminjaman == ""
                if availableBooks kosong then
                    tampilkan "No available books."
                    break
                endif
                tampilkan menu search method
                input(method)
                tampilkan "Enter keyword for Buku:"
                input(keyword)
                if method == "1" then
                    sequentialSearchBukuFiltered(availableBooks, keyword)
                else if method == "2" then
                    selectionSortBuku(availableBooks, true)
                    binarySearchBukuFiltered(availableBooks, keyword)
                else
                    tampilkan "Invalid search method."
                endif
            case "2":
                borrowedBooks <- filter bookData dimana TanggalPeminjaman != ""
                if borrowedBooks kosong then
                    tampilkan "No borrowed books."
                    break
                endif
                tampilkan menu search method
                input(method)
                tampilkan "Enter keyword for Buku:"
                input(keyword)
                if method == "1" then
                    sequentialSearchBukuFiltered(borrowedBooks, keyword)
                else if method == "2" then
                    selectionSortBuku(borrowedBooks, true)
                    binarySearchBukuFiltered(borrowedBooks, keyword)
                else
                    tampilkan "Invalid search method."
                endif
            case "3":
                if studentData kosong then
                    tampilkan "No Mahasiswa data available."
                    break
                endif
                tampilkan "Enter NIM or Name for Mahasiswa:"
                input(keyword)
                found <- false
                for each m in studentData do
                    if keyword ada pada m.NIM atau m.Name then
                        printStudent(m)
                        found <- true
                    endif
                endfor
                if not found then
                    tampilkan "No Mahasiswa found as library member."
                endif
            case "4":
                return
            default:
                tampilkan "Invalid choice."
        endswitch
    endwhile
endprocedure

procedure sequentialSearchBukuFiltered(data : array of Buku, keyword : string)
algoritma:
    foundBooks <- array kosong
    for each b in data do
        if keyword ada pada b.Author atau b.Judul then
            tambah b ke foundBooks
        endif
    endfor
    if foundBooks kosong then
        tampilkan "No Buku data found."
        return
    endif
    tampilkan header tabel buku
    for i = 0 to length(foundBooks)-1 do
        tampilkan data buku ke-i
    endfor
    tampilkan footer tabel buku
endprocedure

procedure binarySearchBukuFiltered(data : array of Buku, keyword : string)
algoritma:
    low <- 0
    high <- length(data)-1
    found <- false
    while low <= high do
        mid <- (low + high) div 2
        if keyword ada pada data[mid].Judul then
            tampilkan header tabel buku
            tampilkan data buku mid
            tampilkan footer tabel buku
            found <- true
            return
        else if keyword < data[mid].Judul then
            high <- mid - 1
        else
            low <- mid + 1
        endif
    endwhile
    if not found then
        tampilkan "No Buku data found."
    endif
endprocedure

procedure printStudent(m : Mahasiswa)
algoritma:
    tampilkan header tabel mahasiswa
    tampilkan data m
    tampilkan footer tabel mahasiswa
endprocedure

procedure printBuku(b : Buku)
algoritma:
    tampilkan header tabel buku
    tampilkan data b
    tampilkan footer tabel buku
endprocedure

procedure printPegawai(p : Pegawai)
algoritma:
    tampilkan header tabel pegawai
    tampilkan data p
    tampilkan footer tabel pegawai
endprocedure

function truncateString(s : string, maxLen : integer) : string
algoritma:
    if panjang(s) <= maxLen then
        return s
    else
        return substring(s, 0, maxLen-3) + "..."
    endif
endfunction