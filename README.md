# FileRef

FileRef is an adapter for gomplate DataSources.
It reads the URL Scheme and provides its content.

## Helmwave

Helmwave uses that library to provide remote StoreFiles and remote Values.


```yaml
files:
    my.yml: # <-- just copy
      src:  https://raw.githubusercontent.com/my.txt
    my.yml.tpl: # <-- render and copy
      src: https://raw.githubusercontent.com/my.txt
    my.yml.sops: # <-- sops and copy
      src: https://raw.githubusercontent.com/my.txt
```

