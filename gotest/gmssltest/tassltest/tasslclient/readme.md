### SM2 ���ܿͻ��� �ⶨ
###  libgmssl.c libgmssl.h��Ϊgolang�ṩ����ʹ��ssl, crypto��C�ӿڿ⣻
```bash
gcc libgmssl.c -I./include -L./ -fPIC -shared -o libgmssl.so
go build main.go
```

###-lssl -lcrypto  ���������Ǵ�TASSL-1.1.1b-V_1.4�����������ġ�
https://github.com/jntass/TASSL-1.1.1b
