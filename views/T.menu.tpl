{{define "menu"}}
<nav class="navbar navbar-inverse">
    <div class="container">
        <a  class="navbar-brand" href="/">
            我的博客
        </a>
        <ui class="nav navbar-nav">
            <li {{if .MyHome}}class="active"{{end}}>
                <a href="/">
                    首页
                </li>
            </a>
            <li {{if .MyCategory}}class="active"{{end}}>
                <a href="/category">
                    分类
                </li>
            </a>
            <li {{if .MyTopic}}class="active"{{end}}>
                <a href="/topic">
                    文章
                </li>
            </a>
        </ui>
        <div class="pull-right">
            <div class="nav navbar-nav ">
                {{if .MyLogin}}
                <li>
                    <a href="/login?exit=true">
                    退出登录
                    </a>
                </li>
                {{else}}
                <li>
                    <a href="/login">
                        管理员登录
                    </a>
                </li>
                {{end}}
            </div>
        </div>
    </div>
</nav>
{{end}}
