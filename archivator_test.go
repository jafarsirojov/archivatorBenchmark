package main

import "testing"

func Benchmark_conArchive(b *testing.B) {
	for i:=0;i<b.N;i++{
		data:=[]string{
			"a.exe",
			"a.txt",
			"b.exe",
			"b.txt",
			"c.exe",
			"c.txt",
			"d.exe",
			"d.txt",
			"f.exe",
			"f.txt",
			"ChromeSetup.exe",
			"go.pdf",
		}
		conArchive(data)
	}
}

func Benchmark_seqArchive(b *testing.B) {
	for i:=0;i<b.N;i++{
		data:=[]string{
			"a.exe",
			"a.txt",
			"b.exe",
			"b.txt",
			"c.exe",
			"c.txt",
			"d.exe",
			"d.txt",
			"f.exe",
			"f.txt",
			"ChromeSetup.exe",
			"go.pdf",
		}
		seqArchive(data)
	}
}