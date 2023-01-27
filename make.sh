echo 生成 libctpquote.so 到 go/lib
cd "$( dirname "${BASH_SOURCE[0]}" )"/go/lib
g++ -shared -fPIC -Wl,-rpath . -o libctpquote.so ../../c/quote.cpp  thostmduserapi_se.so 
