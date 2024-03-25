# env_lsp

Get completions on environment variables based on what is in your `.env` or `.env.local`

## Why
The fact that the only way to access environment variables in go/python is by using plain strings,
thus can't easily get completions really annoyed me and whne I got the chance to be able to something about, I did

Huge thanks to [Tj](https://github.com/tjdevres) on his lsp video, I based the entire thing of his [`educationlsp`](https://github.com/tjdevries/educationalsp)
highly recommend

## supported languages
- [x] Go
- [x] Python
- [ ] Javascript
...

## Usage

1. Install the binary
```sh
go install github.com/musaubrian/env_lsp@latest
```

2. Inform Neovim of the LSP

```lua
local client = vim.lsp.start_client {
  name = "envlsp",
  cmd = { "<path to where it was installed>" },
}

if not client then
  vim.notify "LSP not found"
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = { "go", "python" },
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end
})
```
This tells neovim where to look for the binary, its name and attach it to the current buffer

> [!NOTE]
>
> I should probably improve on this

# caveats
- It depends on the `.env` or `.env.local` to be at the root of the project, anywhere else and it won't work

