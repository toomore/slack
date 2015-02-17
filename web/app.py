# -*- coding: utf-8 -*-
import os
import sys
path = os.path.join('/', *os.path.abspath(__file__).split('/')[:-2])
sys.path.insert(0, path)

import logging
from flask import Flask
from flask import request
from logging.handlers import RotatingFileHandler
from setting import SlackBot_Token
from slack import SlackBot
from grs import RealtimeTWSE
from grs import RealtimeOTC
from grs import RealtimeWeight

app = Flask(__name__)
handler = RotatingFileHandler('./foo.log', maxBytes=10000, backupCount=1)
handler.setLevel(logging.INFO)
app.logger.addHandler(handler)

@app.route("/", methods=['GET', 'POST'])
def main():
    if request.method == "GET":
        return u"Hello World."
    else:
        bot = SlackBot(SlackBot_Token, request.form['team_domain'],
                '#%s' % request.form['channel_name'])
        text = request.form['text']
        if text.startswith('ok stock'):
            result = process_grs(text.replace('ok stock', '').strip())
        else:
            result = text

        result = bot.send(u'@%s %s' % (request.form['user_name'], result))
        app.logger.error(request.form)
        return u''

def process_grs(text):
    if text == 'help':
        return u'ok stock `stock_no` | ok stock weight'

    if text == 'weight':
        weight = RealtimeWeight().data
        return u'加權指數 %s%s, 櫃檯指數 %s%s, 寶島指數 %s%s' % (weight['t00']['price'], weight['t00']['diff'], weight['o00']['price'], weight['o00']['diff'], weight['FRMSA']['price'], weight['FRMSA']['diff'])

    if text != 'help' or text != 'weight':
        stock = RealtimeTWSE(text)
        if stock.data:
            stock = stock.data[text]
        elif not stock.data:
            stock = RealtimeOTC(text)
            if stock.data:
                stock = stock.data[text]
            else:
                return u'查無股票'

        result = u'%s %s $%s %s' % (stock['info']['name'], stock['info']['no'], stock['price'], stock['volume_acc'])

    return result


if __name__ == "__main__":
    app.run()
