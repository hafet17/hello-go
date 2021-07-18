`Definisi Defer`
[comment]: <> (defer func adalah func yang dijadwalkan untuk di eksekusi setelah sebuah func selesai di eksekusi)
[comment]: <> (defer func akan selalu di eksekusi walaupun terjadi error di func yang dieksekusi)

`Definisi Panic`
[comment]: <> (panic func adalah sebuah panic func yang digunakan untuk menghentikan program)
[comment]: <> (panic func biasanya dipanggil ketika terjadi error pada saat program berjalan)
[comment]: <> (saat panic func dipanggil, program akan berhenti, tapi defer akan tetap dieksekusi)

`Definisi Recover`
[comment]: <> (recover adalah sebuah func untuk menangkap data panic)
[comment]: <> (dengan recover proses panic akan terhenti,sehingga progam akan tetap berjalan)