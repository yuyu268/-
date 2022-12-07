#設問１
print('監視ログファイルを入力してください。')
log = list(map(str, input().split()))
#ここで入力をし、入力されたログデータを空白で区切り行ごとにlogへリストとして入れる。
w = '-'
prefin = []
mihu = []
sta = '0'
fin = '0'

for i in range(len(log)):
  if w in log[i] :
    log2 = list(map(str, log[i].split(',')))
#ログデータの中に”-“があるかどうか判定していき、あった場合はその行をさらにカンマで区切り、log2の中へリストとしていれる。
#このとき、log[0]は確認日時、log[1]はサーバアドレス、log[2]は応答時間となる。

    add = log2[1]
    sta = log2[0]
    num = i + 1
#addに応答がなかったサーバアドレスを代入し、staへその確認日時を代入。
#numに応答がなかったサーバアドレスの次のログのリストの位置を代入。

    Y1 = sta[:4]
    M1 = sta[4:6]
    D1 = sta[6:8]
    h1 = sta[8:10]
    m1 = sta[10:12]
    s1 = sta[12:]
#staに代入した最初に応答がなかった日時を、年、月、日、時、分、秒に分けてそれぞれに代入。
    
    for i in range(num, len(log)):
      if add in log[i] :
#応答がなかったサーバアドレスの次のログから、そのサーバアドレスのログを探し出す。

        if w not in log[i]:
          log3 = list(map(str, log[i].split(',')))
          fin = log3[0]
#応答があった場合finへその確認日時を代入。
          if str(fin) not in prefin:
            Y2 = fin[:4]
            M2 = fin[4:6]
            D2 = fin[6:8]
            h2 = fin[8:10]
            m2 = fin[10:12]
            s2 = fin[12:]

            prefin.append(fin)
#複数回応答がなかった際に重複して出力しないよう、prefinに前に処理した、応答があったときの確認日時を保存していき、finと重複していないかを確認。
#その後、重複がなければ応答があった日時を、年、月、日、時、分、秒に分けてそれぞれに代入。
            print('サーバアドレス：', add)
            print('故障期間：', Y1, '年', M1, '月', D1, '日', h1, '時', m1, '分', s1, '秒　～　', Y2, '年', M2, '月', D2, '日', h2, '時', m2, '分', s2, '秒')
            print()

            break
#応答がなかったサーバーアドレスと、応答がなかった確認日時～応答があった確認日時を出力し、処理を終了。

      elif i == len(log)-1 and add not in mihu:
          print('サーバアドレス：', add)
          print('故障期間：', Y1, '年', M1, '月', D1, '日', h1, '時', m1, '分', s1, '秒　～　 未復旧')
          print()
          mihu.append(add)
          break
#ログの終わりまで応答がなかった場合の出力。出力が重複しないよう、一度この処理を行ったサーバアドレスをmihuに記憶し、判定している。
#設問２
print('監視ログファイルを入力してください。')
log = list(map(str, input().split()))
w = '-'
prefin = []
mihu = []
sta = '0'
fin = '0'
N = 2
#Nは回数のパラメータ
for i in range(len(log)):
  if w in log[i] :
    log2 = list(map(str, log[i].split(',')))
    add = log2[1]
    sta = log2[0]
    num = i + 1
    
    Y1 = sta[:4]
    M1 = sta[4:6]
    D1 = sta[6:8]
    h1 = sta[8:10]
    m1 = sta[10:12]
    s1 = sta[12:]

    u = 1
    

    for i in range(num, len(log)):
      if add in log[i] :
        if w in log[i]:
          u = u + 1
#応答がなかったサーバアドレスの次のログから、そのサーバアドレスのログを探し出し、応答がなかった場合はuでカウントする。
        elif w not in log[i] and u >= N:
          log3 = list(map(str, log[i].split(',')))
          fin = log3[0]

          if str(fin) not in prefin:
            Y2 = fin[:4]
            M2 = fin[4:6]
            D2 = fin[6:8]
            h2 = fin[8:10]
            m2 = fin[10:12]
            s2 = fin[12:]

            prefin.append(fin)

            print('サーバアドレス：', add)
            print('故障期間：', Y1, '年', M1, '月', D1, '日', h1, '時', m1, '分', s1, '秒　～　', Y2, '年', M2, '月', D2, '日', h2, '時', m2, '分', s2, '秒')
            print()

            break
#応答があった、かつ、応答がなかった回数がN回以上だった場合、故障期間を出力して終了。

      elif i == len(log)-1 and u >= N and add not in mihu:
          print('サーバアドレス：', add)
          print('故障期間：', Y1, '年', M1, '月', D1, '日', h1, '時', m1, '分', s1, '秒　～　 未復旧')
          print()
          mihu.append(add)
          break
#ログの終わりまで応答がなく、応答がなかった回数がN回以上の場合の出力。出力が重複しないよう、一度この処理を行ったサーバアドレスをmihuに記憶し、判定している。
 

