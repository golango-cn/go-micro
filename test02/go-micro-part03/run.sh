

micro --registry=etcd --api_namespace=mu.micro.book.web  api --handler=web > logs/micro.log &

micro --registry=etcd get service mu.micro.book.srv.auth | grep mu.micro.book.srv.auth-|sed 's/:/ /g'|awk -F ' ' '{print ":"$3}'|lsof -i|grep main|awk -F ' ' '{print $2}'|xargs kill -9
micro --registry=etcd get service mu.micro.book.web.employee | grep mu.micro.book.web.employee-|sed 's/:/ /g'|awk -F ' ' '{print ":"$3}'|lsof -i|grep main|awk -F ' ' '{print $2}'|xargs kill -9
micro --registry=etcd get mu.micro.book.srv.inventory | grep mu.micro.book.srv.inventory-|sed 's/:/ /g'|awk -F ' ' '{print ":"$3}'|lsof -i|grep main|awk -F ' ' '{print $2}'|xargs kill -9
micro --registry=etcd get service mu.micro.book.srv.orders | grep mu.micro.book.srv.orders-|sed 's/:/ /g'|awk -F ' ' '{print ":"$3}'|lsof -i|grep main|awk -F ' ' '{print $2}'|xargs kill -9
micro --registry=etcd get service mu.micro.book.web.orders | grep mu.micro.book.web.orders-|sed 's/:/ /g'|awk -F ' ' '{print ":"$3}'|lsof -i|grep main|awk -F ' ' '{print $2}'|xargs kill -9
micro --registry=etcd get service mu.micro.book.srv.payment | grep mu.micro.book.srv.payment-|sed 's/:/ /g'|awk -F ' ' '{print ":"$3}'|lsof -i|grep main|awk -F ' ' '{print $2}'|xargs kill -9
micro --registry=etcd get service mu.micro.book.web.payment | grep mu.micro.book.web.payment-|sed 's/:/ /g'|awk -F ' ' '{print ":"$3}'|lsof -i|grep main|awk -F ' ' '{print $2}'|xargs kill -9
micro --registry=etcd get service mu.micro.book.srv.employee | grep mu.micro.book.srv.employee-|sed 's/:/ /g'|awk -F ' ' '{print ":"$3}'|lsof -i|grep main|awk -F ' ' '{print $2}'|xargs kill -9

mkdir -p logs

cd auth && go run main.go > ../logs/auth.log &
cd inventory-srv && go run main.go > ../logs/inventory-srv.log &
cd order-srv && go run main.go > ../logs/order-srv.log &
cd order-web && go run main.go > ../logs/order-web.log &
cd payment-srv && go run main.go > ../logs/payment-srv.log &
cd payment-web && go run main.go > ../logs/payment-web.log &

cd svr && go run main.go > ../logs/srv.log &
cd web && go run main.go > ../logs/web.log &

