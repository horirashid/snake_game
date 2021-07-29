git config --global user.name "M212335_wan" # コミットログに残る名前
git config --global user.email "wakeykang@gmail.com" # コミットログに残るメールアドレス
git config --global core.pager "" # ページャーを「なし」．つまり標準出力を選択（標準は less）
git add .
echo 准备提交
git commit -m "岳宇翔"
git push -u origin develop