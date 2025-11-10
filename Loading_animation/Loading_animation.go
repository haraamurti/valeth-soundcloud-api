package Loading_animation

import (
	"context"
	"log"
	"time"
)

func ShowLoadingAnimation(ctx context.Context) {
    // Teks awal
    dots := ".........."
    
    // Loop selamanya sampai ada sinyal "stop"
    for {
        select {
        case <-ctx.Done():
            // Sinyal "stop" (dari 'cancel()') diterima.
            // Kita 'return' untuk menghentikan goroutine ini.
            return
        case <-time.After(800 * time.Millisecond):
            // Belum ada sinyal stop, tunggu 800ms
            // Print log-nya
            log.Printf("loading %s", dots)
            // Tambah titiknya biar makin panjang
            dots = dots + "......."
        }
    }
}