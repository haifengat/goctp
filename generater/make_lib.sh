echo 生成 libctpquote.so 到 go/lib
cd "$( dirname "${BASH_SOURCE[0]}" )"/../lib

g++ -shared -fPIC -Wl,-rpath . -o libctpquote.so ../c/quote.cpp  thostmduserapi_se.so 
g++ -shared -fPIC -Wl,-rpath . -o libctptrade.so ../c/trade.cpp  thosttraderapi_se.so 
