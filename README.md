# elviraBot [Long Term Beta Version]

Bot do HackerSpace Garagem Hacker para informar status e outros dados.

## Gerar Certificados

Altere as opções de acordo com seu dominio.

`
openssl req -newkey rsa:2048 -sha256 -nodes -keyout key.pem -x509 -days 365 -out cert.pem -subj "/C=BR/ST=Paraná/L=Curitiba/O=Garagem Hacker/CN=dominio.com"
`

## Exemplo de como enviar certificado para o telegram sobre seu webhook

__Essa etapa pode ser feita de outra forma pelas ultmias atualizações da apibot do telegram. Tem que melhorar.__

`url -v -F certificate=@/home/usuario/cert.pem -F url=https://dominio.com/bot/ https://api.telegram.org/botTOKEN_DO_SEU_BOT/setWebhook`
