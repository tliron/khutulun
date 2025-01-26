[workspace]
#members = ["crates/cli"]
resolver = "3"

[patch.crates-io]
compris = { path = "../rust-compris/crates/library" }
floria = { path = "../rust-floria/crates/library" }
kutil-cli = { path = "../rust-kutil/cli" }
kutil-cli-macros = { path = "../rust-kutil/cli-macros" }
kutil-std = { path = "../rust-kutil/std" }

[profile.release]
# Especially important for wasm!
strip = "debuginfo"
lto = "thin"        # true is *very* slow to build!
