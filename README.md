# env_lsp

Get completions on environment variables based on what is in your `.env` or `.env.local`

## Why
The fact that the only way to access environment variables in go/python is by using plain strings,
and I couldn't get any completions really annoyed me and when I got the chance to be able to something about, I did

Huge thanks to [Tj](https://github.com/tjdevres) on his lsp video, I based the entire thing of his [`educationalsp`](https://github.com/tjdevries/educationalsp),
highly recommend

## supported languages
- [x] Go
- [ ] Python
- [ ] Javascript
...

## Usage

1. Install the binary
```sh
go install github.com/musaubrian/env_lsp@latest
```

2. Inform Neovim of the LSP

Add this to any file that gets loaded by neovim

```lua
local client = vim.lsp.start_client {
  name = "envlsp",
  cmd = { "<path to where the binary was installed>" }, --I'd recommend the full path
}

if not client then
  vim.notify "LSP not found"
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = { "go"},
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end
})
```
This tells neovim where to look for the binary, its name and attach it to the current buffer

> [!NOTE]
>
> 1. It depends on the `.env` or `.env.local` to be at the root of the project, anywhere else and it won't work
> 2. If both `.env` and `.env.local` exist at the root, it uses the `.env`


