package main

import (
	"errors"
	"fmt"
	"gen/assets"
	"gen/examples/bas_con_mul"
	"gen/examples/bas_con_sin"
	"gen/examples/lay_anc_pad_bot"
	"gen/examples/lay_anc_pad_lef"
	"gen/examples/lay_anc_pad_rig"
	"gen/examples/lay_anc_pad_top"
	"gen/examples/lay_anc_pos_cenxcen"
	"gen/examples/lay_anc_pos_cenxend"
	"gen/examples/lay_anc_pos_cenxsta"
	"gen/examples/lay_anc_pos_endxcen"
	"gen/examples/lay_anc_pos_endxend"
	"gen/examples/lay_anc_pos_endxsta"
	"gen/examples/lay_anc_pos_staxcen"
	"gen/examples/lay_anc_pos_staxend"
	"gen/examples/lay_anc_pos_staxsta"
	"gen/examples/lay_anc_pre"
	"gen/examples/lay_gri_col_1"
	"gen/examples/lay_gri_col_2"
	"gen/examples/lay_gri_col_3"
	"gen/examples/lay_gri_col_4"
	"gen/examples/lay_gri_col_5"
	"gen/examples/lay_gri_dtr_fal_fal"
	"gen/examples/lay_gri_dtr_fal_tru"
	"gen/examples/lay_gri_dtr_tru_fal"
	"gen/examples/lay_gri_dtr_tru_tru"
	"gen/examples/lay_gri_max_all"
	"gen/examples/lay_gri_max_hei"
	"gen/examples/lay_gri_max_off"
	"gen/examples/lay_gri_max_wid"
	"gen/examples/lay_gri_pad_bot"
	"gen/examples/lay_gri_pad_lef"
	"gen/examples/lay_gri_pad_rig"
	"gen/examples/lay_gri_pad_top"
	"gen/examples/lay_gri_pos_cenxcen"
	"gen/examples/lay_gri_pos_cenxend"
	"gen/examples/lay_gri_pos_cenxsta"
	"gen/examples/lay_gri_pos_endxcen"
	"gen/examples/lay_gri_pos_endxend"
	"gen/examples/lay_gri_pos_endxsta"
	"gen/examples/lay_gri_pos_staxcen"
	"gen/examples/lay_gri_pos_staxend"
	"gen/examples/lay_gri_pos_staxsta"
	"gen/examples/lay_gri_pre"
	"gen/examples/lay_gri_spa_0x0"
	"gen/examples/lay_gri_spa_0x25"
	"gen/examples/lay_gri_spa_25x0"
	"gen/examples/lay_gri_spa_25x25"
	"gen/examples/lay_gri_str_1x1"
	"gen/examples/lay_gri_str_234x234"
	"gen/examples/lay_gri_str_234x3"
	"gen/examples/lay_gri_str_3x234"
	"gen/examples/lay_gri_str_3x3"
	"gen/examples/lay_gri_str_3xall"
	"gen/examples/lay_gri_str_5x5"
	"gen/examples/lay_gri_str_allx3"
	"gen/examples/lay_gri_str_allxall"
	"gen/examples/lay_gri_str_allxoff"
	"gen/examples/lay_gri_str_offxall"
	"gen/examples/lay_gri_str_offxoff"
	"gen/examples/lay_row_dir_hor"
	"gen/examples/lay_row_dir_ver"
	"gen/examples/lay_row_max_all"
	"gen/examples/lay_row_max_hei"
	"gen/examples/lay_row_max_off"
	"gen/examples/lay_row_max_wid"
	"gen/examples/lay_row_pad_bot"
	"gen/examples/lay_row_pad_lef"
	"gen/examples/lay_row_pad_rig"
	"gen/examples/lay_row_pad_top"
	"gen/examples/lay_row_pos_cen"
	"gen/examples/lay_row_pos_end"
	"gen/examples/lay_row_pos_sta"
	"gen/examples/lay_row_pre"
	"gen/examples/lay_row_spa_0"
	"gen/examples/lay_row_spa_25"
	"gen/examples/lay_row_spa_50"
	"gen/examples/lay_row_spa_75"
	"gen/examples/lay_row_str_fal"
	"gen/examples/lay_row_str_tru"
	"gen/examples/wid_but_pre"
	"image/png"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type State int

const (
	Switched State = iota
	Updated
)

type Game struct {
	mu        sync.Mutex
	scenes    []Scene
	state     State
	current   int
	offscreen *ebiten.Image
}

func NewGame() *Game {
	root := assets.Output
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == root {
			return nil
		}
		name := d.Name()
		if len(name) > 0 && name[0] == '.' {
			return nil
		}
		return os.Remove(path)
	})

	return &Game{
		scenes: []Scene{
			bas_con_sin.NewGame(),
			bas_con_mul.NewGame(),
			lay_anc_pre.NewGame(),
			lay_anc_pad_lef.NewGame(),
			lay_anc_pad_rig.NewGame(),
			lay_anc_pad_top.NewGame(),
			lay_anc_pad_bot.NewGame(),
			lay_anc_pos_staxsta.NewGame(),
			lay_anc_pos_cenxsta.NewGame(),
			lay_anc_pos_endxsta.NewGame(),
			lay_anc_pos_staxcen.NewGame(),
			lay_anc_pos_cenxcen.NewGame(),
			lay_anc_pos_endxcen.NewGame(),
			lay_anc_pos_staxend.NewGame(),
			lay_anc_pos_cenxend.NewGame(),
			lay_anc_pos_endxend.NewGame(),
			lay_row_pre.NewGame(),
			lay_row_dir_hor.NewGame(),
			lay_row_dir_ver.NewGame(),
			lay_row_pad_lef.NewGame(),
			lay_row_pad_rig.NewGame(),
			lay_row_pad_top.NewGame(),
			lay_row_pad_bot.NewGame(),
			lay_row_spa_75.NewGame(),
			lay_row_spa_25.NewGame(),
			lay_row_spa_50.NewGame(),
			lay_row_spa_0.NewGame(),
			lay_row_str_fal.NewGame(),
			lay_row_str_tru.NewGame(),
			lay_row_pos_sta.NewGame(),
			lay_row_pos_cen.NewGame(),
			lay_row_pos_end.NewGame(),
			lay_row_max_off.NewGame(),
			lay_row_max_wid.NewGame(),
			lay_row_max_hei.NewGame(),
			lay_row_max_all.NewGame(),
			lay_gri_pre.NewGame(),
			lay_gri_col_1.NewGame(),
			lay_gri_col_2.NewGame(),
			lay_gri_col_3.NewGame(),
			lay_gri_col_4.NewGame(),
			lay_gri_col_5.NewGame(),
			lay_gri_str_allxall.NewGame(),
			lay_gri_str_allx3.NewGame(),
			lay_gri_str_3xall.NewGame(),
			lay_gri_str_1x1.NewGame(),
			lay_gri_str_3x3.NewGame(),
			lay_gri_str_5x5.NewGame(),
			lay_gri_str_234x3.NewGame(),
			lay_gri_str_3x234.NewGame(),
			lay_gri_str_234x234.NewGame(),
			lay_gri_str_allxoff.NewGame(),
			lay_gri_str_offxall.NewGame(),
			lay_gri_str_offxoff.NewGame(),
			lay_gri_dtr_tru_tru.NewGame(),
			lay_gri_dtr_tru_fal.NewGame(),
			lay_gri_dtr_fal_tru.NewGame(),
			lay_gri_dtr_fal_fal.NewGame(),
			lay_gri_pad_lef.NewGame(),
			lay_gri_pad_rig.NewGame(),
			lay_gri_pad_top.NewGame(),
			lay_gri_pad_bot.NewGame(),
			lay_gri_spa_0x0.NewGame(),
			lay_gri_spa_0x25.NewGame(),
			lay_gri_spa_25x0.NewGame(),
			lay_gri_spa_25x25.NewGame(),
			lay_gri_max_all.NewGame(),
			lay_gri_max_wid.NewGame(),
			lay_gri_max_hei.NewGame(),
			lay_gri_max_off.NewGame(),
			lay_gri_pos_staxsta.NewGame(),
			lay_gri_pos_cenxsta.NewGame(),
			lay_gri_pos_endxsta.NewGame(),
			lay_gri_pos_staxcen.NewGame(),
			lay_gri_pos_cenxcen.NewGame(),
			lay_gri_pos_endxcen.NewGame(),
			lay_gri_pos_staxend.NewGame(),
			lay_gri_pos_cenxend.NewGame(),
			lay_gri_pos_endxend.NewGame(),
			wid_but_pre.NewGame(),
		},
		offscreen: ebiten.NewImage(
			assets.Frame.Bounds().Dx(),
			assets.Frame.Bounds().Dy(),
		),
		state:   Switched,
		current: 0,
	}
}

func (g *Game) Update() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.current == len(g.scenes) {
		return ebiten.Termination
	}

	if g.state == Switched {
		err := g.scenes[g.current].Update()
		if err != nil {
			return err
		}
		g.state = Updated
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.state == Updated {
		g.scenes[g.current].Draw(screen)

		g.SaveImage(screen)
		g.SaveCode()

		g.current++
		g.state = Switched
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

// END root

func (g *Game) SaveImage(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, assets.Pivot)
	g.offscreen.DrawImage(assets.Frame, nil)
	g.offscreen.DrawImage(screen, op)

	name := FindPackage(g.scenes[g.current])
	base := filepath.Base(name)
	path := filepath.Join(assets.Output, base+".png")
	if err := WriteImage(g.offscreen, path); err != nil {
		log.Fatal(err)
	}
	log.Println("saved", path)

	g.offscreen.Clear()
}

func (g *Game) SaveCode() {
	name := FindPackage(g.scenes[g.current])

	in := RemoveRoot(name)
	in = filepath.Join(assets.Parent, in, "main.go")
	source, err := ReadSource(in)
	if err != nil {
		log.Fatal(err)
	}

	source = assets.RegPackage.ReplaceAllString(source, "package main")
	source = assets.RegAssets.ReplaceAllString(source, "")
	source = assets.RegImports.ReplaceAllStringFunc(source, func(s string) string {
		return assets.RegNewline.ReplaceAllString(s, "\n")
	})
	source = assets.RegImports.ReplaceAllStringFunc(source, func(match string) string {
		return match[:len(match)-1] + "\t" + assets.ColorPackage + "\n)"
	})
	source = assets.RegColor.ReplaceAllStringFunc(source, func(match string) string {
		return assets.ColorMap[match]
	})
	matches := assets.RegComment.FindAllStringIndex(source, -1)
	if len(matches) > 0 {
		last := matches[len(matches)-1][0]
		source = source[:last] + assets.PartMain + source[last:]
	}

	out := filepath.Base(name)
	out = filepath.Join(assets.Output, out+".txt")
	err = WriteSource(source, out)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("saved", out)
}

func FindPackage(i any) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.PkgPath()
}

func WriteImage(img *ebiten.Image, path string) (e error) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer Close(&e, f)
	if err = png.Encode(f, img); err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	return nil
}

func ReadSource(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read: %w", err)
	}
	return string(data), err
}

func WriteSource(source, path string) error {
	dirs := filepath.Dir(path)
	if _, err := os.Stat(dirs); os.IsNotExist(err) {
		if err := os.MkdirAll(dirs, 0700); err != nil {
			return fmt.Errorf("mkdir: %w", err)
		}
	}
	err := os.WriteFile(path, []byte(source), 0666)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

func RemoveRoot(path string) string {
	sep := "/"
	slash := filepath.ToSlash(path)
	parts := strings.Split(slash, sep)
	if len(parts) > 1 && parts[0] == "" {
		parts = parts[1:]
	}
	return strings.Join(parts[1:], sep)
}

func Close(dest *error, c io.Closer) {
	*dest = errors.Join(*dest, c.Close())
}

func main() {
	log.SetFlags(0)
	ebiten.SetWindowSize(assets.Width, assets.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
