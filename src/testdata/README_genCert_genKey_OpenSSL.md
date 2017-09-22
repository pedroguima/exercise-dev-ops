## Creating SSL/TLS certificates

So basically, and to make it simple, use this command (details below).

    openssl req -x509 -newkey rsa:4096 -sha256 -nodes -keyout example.key -out example.crt -subj "/CN=x.test.youtube.com" -days 3650

It generates 2 files:

1. a Private Key: *"example.key"*
2. a Public Key or Certificate: *"example.crt"*

The extensions 'key' and 'crt' are conventional.

The Client will have to know the public key: the 'crt'. Look into the Client and Server code, searching for 'crt'.
The Server will have to know both: the 'crt' and the 'key'. Look into the Server code, searching for 'key'.



## Generating strong certificates and keys

The following command creates a relatively strong (as of 2017) certificate for the domain example.com that is valid for 3650 days (~10 years).
It saves the private key and certificate into the files example.key and example.crt.

	openssl req -x509 -newkey rsa:4096 -sha256 -nodes -keyout example.key -out example.crt -subj "/CN=x.test.youtube.com" -days 3650

Since it is a self-signed certificate that has to be accepted by users manually, it doesn't make sense to use a short expiration or weak crypto.

In the future, you might want to use more than 4096 bits for the RSA key and a hash algorithm stronger than sha256, but as of 2017 these are sane values.
They are sufficiently strong while being supported by all modern browsers.

Side note: Theoretically you could leave out the -nodes parameter (which means "no DES encryption"), in which case example.key would be encrypted with a password.
However, this is almost never useful for a server installation, because it means that you either have to store the password on the server as well,
or that you have to enter it manually on each reboot.


DESCRIPTION:

The OPENSSL documentatiion is here: https://www.openssl.org/docs/manmaster/apps/req.html


    req

    PKCS#10 certificate request and certificate generating utility.

    -x509

    this option outputs a self signed certificate instead of a certificate request. This is typically used to generate a test certificate or a self signed root CA.

    -newkey arg

    this option creates a new certificate request and a new private key. The argument takes one of several forms. rsa:nbits, where nbits is the number of bits, generates an RSA key nbits in size.

    -keyout filename

    this gives the filename to write the newly created private key to.

    -out filename

    This specifies the output filename to write to or standard output by default.

    -days n

    when the -x509 option is being used this specifies the number of days to certify the certificate for. The default is 30 days.

    -nodes

    if this option is specified then if a private key is created it will not be encrypted.

The documentation is actually more detailed than the above, I just summarized it here.
