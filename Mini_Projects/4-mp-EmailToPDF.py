import imaplib
import base64
import email
from typing import Text
import os

from fpdf import FPDF
import traceback
import stdiomask

USER_EMAIL = input('User Email: ')

FROM_PWD = stdiomask.getpass() 
FROM_MAIL = input('From Mail : ')
# FROM_EMAIL = 'reddyanil1445@gmail.com'
# FROM_PWD = 'password'
SMTP_SERVER = "imap.gmail.com"
SMTP_PORT = 993


def get_dody(raw_data):
    if raw_data.is_multipart():
        # print(raw_data.get_payload(0))
        return get_dody(raw_data.get_payload(0))
    else:
        return raw_data.get_payload(None,True) 

def get_attachment_files(mail):
    folder_name = 'attachments/'
    for part in mail.walk():

            if(part.get('Content-Disposition') is None):
                continue
            file_name = part.get_filename()
            if(bool(file_name)):
                file_path = os.path.join(folder_name, file_name)
                with open(file_path, 'wb') as file:
                    file.write(part.get_payload(decode=True))
                

def read_email_from_gmail():
    try:
        mail_con = imaplib.IMAP4_SSL(SMTP_SERVER,SMTP_PORT)
        mail_con.login(USER_EMAIL,FROM_PWD)

        mail_con.select("inbox")
        # data = mail_con.search(None,'ALL')
        searchString = ''
        if FROM_MAIL != '' :
            searchString = '(FROM "'+FROM_MAIL+'")'
        else:
            searchString='ALL'    
        data = mail_con.search(None,searchString)
        print(len(data))
      
        mailids = data[1]
        id_list = mailids[0].split()
        print(id_list)
        if len(id_list) == 0 :
            print("No Mail Found ..")
            return
        first_email_id = id_list[0]
        last_email_id = id_list[-1]
        
        # reading last mail from inbox
        result , data = mail_con.fetch(last_email_id,'(RFC822)')
        print(result)
        raw_data = email.message_from_bytes(data[0][1])
        sub = raw_data.get("SUBJECT")
        from_mail = raw_data.get("FROM")
        to_mail = raw_data.get("TO")
        date = raw_data.get('DATE')
        print(date)
        print(sub)
        print(from_mail)

        pdf = FPDF()

        pdf.add_page()
        pdf.set_font("Arial", size = 15)
        pdf.cell(150,5,txt="DATE : "+date, ln=True, border=True)
        pdf.cell(150,5,txt="FROM : "+from_mail, ln=True, border=True)
        pdf.cell(150,5,txt="TO :  "+to_mail, ln=2, border=True)
        pdf.cell(150,5,txt="Subject :  "+sub, ln=2, border=True)
        pdf.cell(5,5,txt='',ln=True)

        # print(raw.get("FROM"))
        # get_attachment_files(raw_data)
        # pdf.add_font('DejaVu', '', 'DejaVuSansCondensed.ttf', uni=True)
        folder_name = 'attachments/'
        for part in raw_data.walk():
            if part.get_content_maintype() == 'multipart':
                continue
            if(part.get('Content-Disposition') is None):
                continue
            file_name = part.get_filename()
            if(bool(file_name)):
                file_path = os.path.join(folder_name, file_name)
                with open(file_path, 'wb') as file:
                    file.write(part.get_payload(decode=True))
                pdf.cell(150,5,txt=file_name,ln=True,link= os.path.abspath(file_path), border=True)
                print(file_path)
        
        body = get_dody(raw_data)
        body = body.decode('UTF-8')
        print(body)   

        print(os.path.abspath(file_path))
        
        pdf.cell(5,5,txt='',ln=True)
        pdf.multi_cell(0, 5, txt=body) 
        # pdf.multi_cell(0,5,txt="Atachhment : "+os.path.abspath(file_path))

        pdf.output("email.pdf")

        mail_con.close()
        mail_con.logout()
    except Exception as e:
        print(e)
        if(mail_con is not None):
            mail_con.logout()
        pass 
        




read_email_from_gmail()

