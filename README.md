[![GhostWardrobe.nar](https://img.shields.io/github/v/release/apxxxxxxe/GhostWardrobe?color=%2367ab7b&label=GhostWardrobe.nar&logo=github)](https://github.com/apxxxxxxe/GhostWardrobe/releases/latest/download/GhostWardrobe.nar) 
[![commits](https://img.shields.io/github/last-commit/apxxxxxxe/GhostWardrobe?color=%2367ab7b&label=%E6%9C%80%E7%B5%82%E6%9B%B4%E6%96%B0&logo=github)](https://github.com/apxxxxxxe/GhostWardrobe/commits/main)
[![commits](https://img.shields.io/tokei/lines/github/apxxxxxxe/GhostWardrobe?color=%2367ab7b)](https://github.com/apxxxxxxe/GhostWardrobe/commits/main)

# 伺かプラグイン「GhostWardrobe」

![ghostwardrobe_sample](https://user-images.githubusercontent.com/39634779/230760848-df3d9331-0688-4f00-a10d-bf35d0349a4f.gif)

- SSPでのみ動作確認
- このプラグインは、YAYA as PLUGINを使用して作られています。
- 作成にあたり、プラグイン「第弐版仮想道頓堀水泳拡張」をテンプレートとして使用させていただきました。

## 何をするもの？
任意のゴーストの着せ替えを保存し、いつでも呼び出せるようにします。  
複数の着せ替えパターンを保存することが可能です。

## どうやって使うの？
右クリックメニューのプラグインから実行するとメニューが開き、「保存」「読み込み」が可能です。

## インストール方法
ゴーストのインストールと同様に、本プラグインのnarファイルを起動中のゴーストにドラッグ＆ドロップしてください。  

## 注意
インストール直後はバージョンが古い場合があるため、必ずネットワーク更新を行ってください。  
本プラグインの右クリックメニューからネットワーク更新が可能です。

## 現在実装済みの機能
- 各ゴーストの着せ替え保存/読み込み
  - 各キャラクター(\0,\1,...)に対応
  - 各シェルに対応

## 着せ替え保存データの名前について
各保存データの名前は英数字の組み合わせですが、これは保存データの実体である着せ替えスクリプト(\\![bind,...)をCRC32でハッシュ化したものです。  
同じ着せ替えの組み合わせをハッシュ化した場合、同じデータ名が得られます。  
よって同じ内容のデータが複数保存されることはありません。

## ダウンロード
[![GhostWardrobe.nar](https://img.shields.io/github/v/release/apxxxxxxe/GhostWardrobe?color=%2367ab7b&label=GhostWardrobe.nar&logo=github)](https://github.com/apxxxxxxe/GhostWardrobe/releases/latest/download/GhostWardrobe.nar) 
