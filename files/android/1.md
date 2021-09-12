## 优点
 mvvm的优点是代码耦合度低，方便进行单元测试，viewmodel可以复用。
## 使用
1.集成databinding

```java
android {
    compileSdkVersion 25
    buildToolsVersion "25.0.2"
    ...
    dataBinding {
        enabled true//添加这个
    }
}
```
2.在xml里面

```xml
<?xml version="1.0" encoding="utf-8"?>
<layout xmlns:android="http://schemas.android.com/apk/res/android">

    <data>

        <import type="com.haijia.mvvp.bean.User" />

        <variable
            name="user"
            type="User" />

    </data>

    <LinearLayout
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:orientation="vertical">

        <TextView
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:text="@{String.valueOf(user.id)}" />

        <TextView
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:text="@{user.realName}" />


    </LinearLayout>

</layout>

```
3.在activity里

```java
public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        ActivityMainBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_main);
        binding.setUser(new User(123,"a1"));
    }
}
```
## 原理
1.DataBindingUtil.setContentView方法将xml中的各个View赋值给ViewDataBinding，完成findviewbyid的任务
2.ViewDataBinding的setVariable方法建立了ViewDataBinding与VM之间的联系，也就搭建了一个可以互相通信的桥梁
3.当VM层调用notifyPropertyChanged方法时，最终在ViewDataBinding的executeBindings方法中处理逻辑



