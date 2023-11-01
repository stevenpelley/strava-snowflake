import json

from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives.asymmetric import dsa
from cryptography.hazmat.primitives import serialization

from snowflake.snowpark import Session

def create_session():
    # run from project root
    with open("../snowflake_config.json", "r") as f:
        j = json.loads(f.read())

    with open(j['private_key_file'], "rb") as key:
        p_key= serialization.load_pem_private_key(
            key.read(),
            password=None,
            backend=default_backend()
        )

    pkb = p_key.private_bytes(
        encoding=serialization.Encoding.DER,
        format=serialization.PrivateFormat.PKCS8,
        encryption_algorithm=serialization.NoEncryption())

    sess = Session.builder.configs(
        {
            "user":j['user'],
            "account":j['account'],
            "private_key":pkb,
        }).create()
    sess.add_import('src/')
    sess.add_packages('snowflake-snowpark-python')
    return sess