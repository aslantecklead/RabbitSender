-- Создаем таблицу Sender
CREATE TABLE sender (
                        ID_Sender serial PRIMARY KEY,
                        Sender_mail text NULL
);

-- Создаем таблицу EmailMessage
CREATE TABLE emailmessage (
                              id_email_message serial PRIMARY KEY,
                              email_title text NULL,
                              email_body text NULL,
                              receiver_mail text NULL,
                              ID_Sender int NULL,
                              FOREIGN KEY (ID_Sender) REFERENCES Sender(ID_Sender)
);

-- Создаем таблицу MessagesLogs
CREATE TABLE messageslogs (
                              ID_MessagesLogs serial PRIMARY KEY,
                              timestamp timestamp NULL,
                              error_message text NULL,
                              id_email_message int NULL,
                              FOREIGN KEY (id_email_message) REFERENCES EmailMessage(id_email_message)
);