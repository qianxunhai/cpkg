package main

import (
    "flag"
    "fmt"
    "github.com/xlab/treeprint"
    tpl "github.com/qianxunhai/cpkg/template"
    "os"
    "path"
    "path/filepath"
    "strings"
    "text/template"
    "time"
)

type config struct {
    // foo
    Name string
    //go build   /usr/local/go1.14/bin
    GoRoot string
    // micro new example -type
    FQDN string
    // github.com/micro/foo
    Dir string
    // $GOPATH/src/github.com/micro/foo
    GoDir string
    // $GOPATH
    GoPath string
    // UseGoPath
    UseGoPath bool
    // Files
    Files []file
    // Comments
    Comments []string
    // Plugins registry=etcd:broker=nats
    Plugins []string
}
type file struct {
    Path string
    Tmpl string
}

func create(c config) error {
    // check if dir exists
    if _, err := os.Stat(c.GoDir); !os.IsNotExist(err) {
        return fmt.Errorf("%s already exists", c.GoDir)
    }

    // just wait
    <-time.After(time.Millisecond * 250)

    fmt.Printf("Creating service %s in %s\n\n", c.FQDN, c.GoDir)

    t := treeprint.New()

    nodes := map[string]treeprint.Tree{}
    nodes[c.GoDir] = t

    // write the files
    for _, file := range c.Files {
        f := filepath.Join(c.GoDir, file.Path)
        dir := filepath.Dir(f)

        b, ok := nodes[dir]
        if !ok {
            d, _ := filepath.Rel(c.GoDir, dir)
            b = t.AddBranch(d)
            nodes[dir] = b
        }

        if _, err := os.Stat(dir); os.IsNotExist(err) {
            if err := os.MkdirAll(dir, 0755); err != nil {
                return err
            }
        }

        p := filepath.Base(f)

        b.AddNode(p)
        if err := write(c, f, file.Tmpl); err != nil {
            return err
        }
    }

    // print tree
    fmt.Println(t.String())

    for _, comment := range c.Comments {
        fmt.Println(comment)
    }

    // just wait
    <-time.After(time.Millisecond * 250)

    return nil
}

func main(){
    fmt.Println()
    useGoModule := os.Getenv("GO111MODULE")
    var dir,name,goRoot string
    flag.StringVar(&goRoot, "goroot", "/usr/local/go1.14/bin", "构建时 go目录 /usr/local/go1.14/bin")
    flag.StringVar(&dir, "dir", "", "创建目录")
    name = os.Args[1]
    flag.Parse()

    if name=="" {
        fmt.Println("项目名不能为空")
       return
    }

    c := config{
        Name:      name,
        Dir:       dir,
        GoRoot:    goRoot,
        Files: []file{
            {name+"/main.go", tpl.MainFunc},
            {name+"/go-build.sh", tpl.Sh_go_build},
            {name+"/pkg/gmysql/mysql.go", tpl.Mysql},
            {name+"/pkg/gredis/cmd.go", tpl.Redis_cmd},
            {name+"/pkg/gredis/redis.go", tpl.Redis},
            {name+"/pkg/logging/core.go", tpl.Log_core},
            {name+"/pkg/logging/log.go", tpl.Log_log},
            {name+"/pkg/setting/setting.go", tpl.Setting},
            {name+"/shell/check_monitor.sh", tpl.Sh_check_monitor},
            {name+"/shell/monitor_exec.sh", tpl.Sh_monitor_exec},
            {name+"/shell/publish-script.sh", tpl.Sh_publish_script},
            {name+"/conf/app.ini", tpl.Conf},
            {name+"/routers/route.go", tpl.Route},
            {name+"/controllers/exampleController/example.go", tpl.ExampleController},
            {name+"/pkg/app/response.go", tpl.ApiResponse},
        },
    }

    if path.IsAbs(dir) {
        fmt.Println("require relative path as service will be installed in GOPATH")
        return
    }

    if useGoModule != "off" {
        c.Files = append(c.Files, file{name+"/go.mod", tpl.Module})
    }

    if err := create(c); err != nil {
        fmt.Println(err)
        return
    }
}
func write(c config, file, tmpl string) error {
    fn := template.FuncMap{
        "title": strings.Title,
    }

    f, err := os.Create(file)
    if err != nil {
        return err
    }
    defer f.Close()

    t, err := template.New("f").Funcs(fn).Parse(tmpl)
    if err != nil {
        return err
    }

    return t.Execute(f, c)
}