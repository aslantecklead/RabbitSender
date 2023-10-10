CREATE TABLE IF NOT EXISTS MsgRecived (
                                          FromEmail text NULL,
                                          ToEmail text NULL,
                                          MsgTitle text NULL,
                                          EmailBody text NULL,
                                          Timestamp int NULL
);

CREATE TABLE IF NOT EXISTS Sender (
                                      IDSender serial PRIMARY KEY,
                                      SenderMail text NULL
);

CREATE TABLE IF NOT EXISTS EmailMessage (
                                            IDEmailMessage serial PRIMARY KEY,
                                            EmailTitle text NULL,
                                            EmailBody text NULL,
                                            ReceiverMail text NULL,
                                            IDSender int NULL
);
