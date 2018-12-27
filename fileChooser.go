package faithdroid

import (
	"fmt"
	"io/ioutil"
	"mime"
	"os"
	"strings"
)

type FFileChooser struct {
	a                                               IActivity
	showAfter                                       bool
	title, positiveButtonTitle, negativeButtonTitle string
	onResult                                        func(string)
	onResults                                       func([]string)
	filter, hiddenFilter                            func(string) bool
	multiple, folderMode                            bool

	currentDir string
	flist      []*FFileItem
}

func FileChooser(a IActivity) *FFileChooser {
	f := &FFileChooser{a: a, positiveButtonTitle: "确定", negativeButtonTitle: "取消"}
	f.hiddenFilter = func(f string) bool {
		if f[:1] == "." {
			return false
		}
		return true
	}
	return f
}
func (f *FFileChooser) ButtonTitle(positive, negative string) *FFileChooser {
	f.positiveButtonTitle = positive
	f.negativeButtonTitle = negative
	return f
}
func OpenSystemFileChooser(a IActivity, selectType string, allowMultiple bool, callback func([]string)) {
	if !CheckSelfPermission(a, Permissions.READ_EXTERNAL_STORAGE) {
		RequestPermissions(a, []string{Permissions.READ_EXTERNAL_STORAGE}, func(bs []bool) {
			OpenSystemFileChooser(a, selectType, allowMultiple, callback)
		})
		return
	}
	fnId := NewToken()
	fn := func(s string) string {
		var as []string
		UnJson(s, &as)
		callback(as)
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "OpenFileChooser", JsonArray([]string{
		selectType,
		SPrintf(allowMultiple),
		fnId,
	}))
}
func (f *FFileChooser) Title(s string) *FFileChooser {
	f.title = s
	return f
}

func (f *FFileChooser) StartsFromLocation(dir string) *FFileChooser {
	f.currentDir = dir
	return f
}

// filter:"image/*.png"
func (f *FFileChooser) Filter(s func(fname string) bool) *FFileChooser {
	f.filter = s
	return f
}
func (f *FFileChooser) TypeSelectFile(rp func(string)) *FFileChooser {
	f.onResult = rp
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFiles(rp func([]string)) *FFileChooser {
	f.onResults = rp
	f.multiple = true
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFolder(rp func(string)) *FFileChooser {
	f.folderMode = true
	f.onResult = rp
	if f.showAfter {
		f.Show()
	}
	return f
}

type FFileItem struct {
	checked bool
	info    os.FileInfo
	path    string
	icon    string
}

func getDirOfFile(path string) string {
	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}
	for i := len(path) - 1; i > -1; i-- {
		if path[i:i+1] == "/" {
			return path[:i]
		}
	}
	return path
}
func (f *FFileChooser) Show() {
	StartActivity(f.a, func(a IActivity) {
		if !CheckSelfPermission(a, Permissions.WRITE_EXTERNAL_STORAGE) {
			RequestPermissions(a, []string{Permissions.WRITE_EXTERNAL_STORAGE}, nil)
		}
		loadCurrentDir := func() {
			go func() {
				extSPath := GetExternalStorageDirectory(a)
				if f.currentDir == "" {
					f.currentDir = Getrpath(extSPath)
				}
				simplifiedPath := strings.Replace(f.currentDir, Getrpath(extSPath), "内部存储/", -1)
				infos, e := ioutil.ReadDir(f.currentDir)
				if e != nil {
					fmt.Println(`fileChooser loadCurrentDir error :`, e)
					RunOnUIThread(a, func() {
						ShowToast(a, e.Error())
					})
					return
				}
				f.flist = nil
				var fileList, dirList []*FFileItem
				for _, v := range infos {
					if f.folderMode && !v.IsDir() {
						continue
					}
					if f.filter != nil && !f.filter(v.Name()) {
						continue
					}
					if f.hiddenFilter != nil && !f.hiddenFilter(v.Name()) {
						continue
					}
					ffi := &FFileItem{
						info: v,
						path: f.currentDir + v.Name(),
						icon: getIconURLByFileType(f.currentDir + v.Name()),
					}
					if v.IsDir() {
						dirList = append(dirList, ffi)
					} else {
						fileList = append(fileList, ffi)
					}
				}
				f.flist = append(dirList, fileList...)
				RunOnUIThread(a, func() {
					GetToolbarById("faithdroid/filechooser/toolbar").SubTitle(simplifiedPath)
					GetListViewById("faithdroid/filechooser/pages").NotifyDataSetChanged()
				})
			}()
		}

		SetOnBackPressed(a, func() bool {
			if strings.Count(GetToolbarById("faithdroid/filechooser/toolbar").GetSubTitle(), "/") < 2 {
				return true
			}
			f.currentDir = Getrpath(getDirOfFile(f.currentDir))
			loadCurrentDir()
			return false
		})
		getTwoButtons := func() IView {
			if f.multiple || f.folderMode {
				return LinearLayout(a).Horizontal().Size(-2, -1).PaddingAll(10).Append(
					Button(a).Text(f.positiveButtonTitle).Background(Backgrounds.RadiusCorner).TextColor(Colors.White).Foreground(Colors.RippleEffect).LayoutWeight(1).SetId("faithdroid/filechooser/submit").OnClick(func() {
						if f.multiple {
							var flist []string
							for _, v := range f.flist {
								if v.checked {
									flist = append(flist, v.path)
								}
							}
							f.onResults(flist)
							Finish(a)
						} else {
							if f.folderMode {
								f.onResult(f.currentDir)
								Finish(a)
							} else {
								ShowToast(a, "You're not support to click me")
								Finish(a)
							}
						}
					}),
					Button(a).Text(f.negativeButtonTitle).LayoutWeight(1).BackgroundColor(Colors.Grey).TextColor(Colors.White).Foreground(Colors.RippleEffect).SetId("faithdroid/filechooser/cancel").OnClick(func() {
						Finish(a)
					}),
				)
			}
			return nil
		}
		LinearLayout(a).DeferShow().Size(-2, -2).Append(
			Toolbar(a).SetId("faithdroid/filechooser/toolbar").Title(f.title).Menus(
				MenuItem("反选").OnClick(func() {
					for k, v := range f.flist {
						if v.info.IsDir() {
							continue
						}
						f.flist[k].checked = !f.flist[k].checked
					}
					GetListViewById("faithdroid/filechooser/pages").NotifyDataSetChanged()
				}),
				MenuItem("隐藏文件").OnClick(func() {
					if f.hiddenFilter == nil {
						f.hiddenFilter = func(f string) bool {
							if f[:1] == "." {
								return false
							}
							return true
						}
					} else {
						f.hiddenFilter = nil
					}
					loadCurrentDir()
				}),
			),
			VListView(
				a,
				func(lv *FListView) IView {
					return LinearLayout(a).Horizontal().Size(-2, -1).Foreground(Colors.RippleEffect).PaddingAll(10).Append(
						CheckBox(a).SetItemId(lv, "check").LayoutGravity(Gravitys.CenterVertical),
						ImageView(a).SetItemId(lv, "icon").LayoutGravity(Gravitys.CenterVertical).Size(60, 60).MarginRight(20).ScaleType(ScaleTypes.CenterInside),
						TextView(a).SetItemId(lv, "text").LayoutGravity(Gravitys.CenterVertical),
					).SetItemId(lv, "ctn")
				},
				func(vh *ViewHolder, pos int) {
					check := vh.GetCheckBoxByItemId("check")
					check.Checked(f.flist[pos].checked)
					if !f.multiple {
						check.Gone()
					} else {
						if f.folderMode {
							if f.flist[pos].info.IsDir() {
								check.Visible()
								check.Enabled(true)
							}
						} else {
							if f.flist[pos].info.IsDir() {
								check.Invisible()
								check.Enabled(false)
							} else {
								check.Visible()
								check.Enabled(true)
							}
						}
					}
					vh.GetImageViewByItemId("icon").CachedSrc(f.flist[pos].icon)
					vh.GetTextViewByItemId("text").Text(f.flist[pos].info.Name())
				},
				func() int {
					return len(f.flist)
				},
			).OnItemClick(func(pos int) {
				if f.flist[pos].info.IsDir() {
					if f.folderMode {
						if f.multiple {
							f.flist[pos].checked = !f.flist[pos].checked
							GetListViewById("faithdroid/filechooser/pages").NotifyDataSetChanged()
						} else {
							f.currentDir = Getrpath(f.currentDir) + f.flist[pos].info.Name() + "/"
							loadCurrentDir()
						}
					} else {
						f.currentDir = Getrpath(f.currentDir) + f.flist[pos].info.Name() + "/"
						loadCurrentDir()
					}
				} else {
					if !f.folderMode {
						if f.multiple {
							f.flist[pos].checked = !f.flist[pos].checked
							GetListViewById("faithdroid/filechooser/pages").NotifyDataSetChanged()
						} else {
							f.onResult(f.flist[pos].path)
							Finish(a)
						}
					}
				}
			}).Size(-2, -2).LayoutWeight(1).SetId("faithdroid/filechooser/pages"),
			getTwoButtons(),
		)
		loadCurrentDir()
	}, nil)
}
func (f *FFileChooser) DeferShow() *FFileChooser {
	f.showAfter = true
	return f
}
func getIconURLByFileType(fpath string) string {
	f, e := os.Open(fpath)
	if e != nil {
		fmt.Println(`os.open error :`, e)
		return ""
	}
	server := "https://jywjl.github.io/images/icons/"
	info, e := f.Stat()
	if e != nil {
		return server + "file.png"
	}
	if info.IsDir() {
		return server + "folder.png"
	}
	nameS := strings.Split(f.Name(), ".")
	ext := nameS[len(nameS)-1]
	mimeTypes := strings.Split(mime.TypeByExtension("."+ext), "/")
	switch mimeTypes[0] {
	case "audio":
		return server + "audio.png"
	case "image":
		return "file://" + fpath
	case "video":
		return server + "video.png"
	default:
		return server + "file.png"
	}
}
