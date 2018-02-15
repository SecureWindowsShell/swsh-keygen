/*
    Copyright (C) 2018  Muhammad Muzzammil
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
    "crypto/rsa"
    "crypto/rand" 
    "encoding/pem" 
    "crypto/x509" 
    "golang.org/x/crypto/ssh" 
    "io/ioutil"
    "os"
    "flag"
)

func main(){
    var pubKeyPath, privateKeyPath string
    flag.StringVar(&pubKeyPath, "pub", "nil", "PublicKey path")
    flag.StringVar(&privateKeyPath, "pri", "nil", "PrivateKey path")
    flag.Parse();
    if pubKeyPath != "nil" && privateKeyPath != "nil" {
        privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
        if err != nil { panic(err) }
        privateKeyFile, err := os.Create(privateKeyPath)
        defer privateKeyFile.Close()
        if err != nil { panic(err) }
        privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
        if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil { panic(err) }
        pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
        if err != nil { panic(err) }
        ioutil.WriteFile(pubKeyPath, ssh.MarshalAuthorizedKey(pub), 0655)
    } else { panic("Values cannot be nil.") }
}
