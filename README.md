# Arcane Web
This webapp is built with .NET Core and Razor Pages.

# How to run in VSCode
```
cd arcane-web
code -r arcane-web
```

# How to build project and open in browser locally
# Open browser with "localhost:[whatever-port-it-gives-you]
```
cd arcane-web
dotnet run
```

# Creating future additional Razor / .NET pages
```
cd [path-to-your-local-repo-folder]
dotnet new webapp -o [project-name]
code -r [project-name]
```

After running this, it will run the HTTPS version on port 5001, and the HTTP version on 5000.