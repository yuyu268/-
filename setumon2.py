print('�Ď����O�t�@�C������͂��Ă��������B')
log = list(map(str, input().split()))
w = '-'
prefin = []
mihu = []
sta = '0'
fin = '0'
N = 2

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


            print('�T�[�o�A�h���X�F', add)
            print('�̏���ԁF', Y1, '�N', M1, '��', D1, '��', h1, '��', m1, '��', s1, '�b�@�`�@', Y2, '�N', M2, '��', D2, '��', h2, '��', m2, '��', s2, '�b')
            print()

            break

      elif i == len(log)-1 and u >= N and add not in mihu:
          print('�T�[�o�A�h���X�F', add)
          print('�̏���ԁF', Y1, '�N', M1, '��', D1, '��', h1, '��', m1, '��', s1, '�b�@�`�@ ������')
          print()
          mihu.append(add)
          break