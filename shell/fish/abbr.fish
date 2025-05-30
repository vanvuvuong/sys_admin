abbr 'drm!' "docker container rm -f"
abbr ainfo "sudo apt list --installed | grep"
abbr arem "sudo apt autoremove"
abbr audioreload "sudo udevadm control --reload-rules && sudo udevadm trigger"
abbr bb "kubectl run busybox --image=busybox:1.28 --rm -it --restart=Never"
abbr c clear
abbr clean "sudo apt clean"
abbr cnf command_not_found_handle
abbr dbl "docker build"
abbr dcb "docker compose build"
abbr dcdn "docker compose down"
abbr dce "docker compose exec"
abbr dcin "docker container inspect"
abbr dck "docker compose kill"
abbr dcl "docker compose logs"
abbr dclF "docker compose logs -f --tail 0"
abbr dclf "docker compose logs -f"
abbr dcls "docker container ls"
abbr dclsa "docker container ls -a"
abbr dco "docker compose"
abbr dcps "docker compose ps"
abbr dcpull "docker compose pull"
abbr dcr "docker compose run"
abbr dcrestart "docker compose restart"
abbr dcrm "docker compose rm"
abbr dcstart "docker compose start"
abbr dcstop "docker compose stop"
abbr dcup "docker compose up"
abbr dcupb "docker compose up --build"
abbr dcupd "docker compose up -d"
abbr dcupdb "docker compose up -d --build"
abbr dib "docker image build"
abbr dii "docker image inspect"
abbr dils "docker image ls"
abbr dipu "docker image push"
abbr dirm "docker image rm"
abbr disservice "sudo systemctl disable"
abbr dit "docker image tag"
abbr dlo "docker container logs"
abbr dnc "docker network create"
abbr dncn "docker network connect"
abbr dndcn "docker network disconnect"
abbr dni "docker network inspect"
abbr dnls "docker network ls"
abbr dnrm "docker network rm"
abbr dpo "docker container port"
abbr dps "docker ps"
abbr dpsa "docker ps -a"
abbr dpu "docker pull"
abbr dr "docker container run"
abbr drit "docker container run -it"
abbr drm "docker container rm"
abbr drs "docker container restart"
abbr dst "docker container start"
abbr dsta "docker stop \$(docker ps -q)"
abbr dstp "docker container stop"
abbr dtop "docker top"
abbr dvi "docker volume inspect"
abbr dvls "docker volume ls"
abbr dvprune "docker volume prune"
abbr dxc "docker container exec"
abbr dxcit "docker container exec -it"
abbr e exit
abbr enservice "sudo systemctl enable"
abbr ffe "find . -type f -exec"
abbr fxs "sudo find / ! -path '/proc/*' -prune ! -path '/sys/*' -prune ! -path '/var/*' -prune ! -path '/run/*' -prune ! -path '/mnt/*' -prune"
abbr idea eureka
abbr ins "sudo apt install"
abbr insdoc "sudo apt list -I | sed -rn '/-doc/! s/([a-z-]+[a-z]).*/\1/p' | sudo xargs -tI§ apt add §-doc"
abbr ipip "pip install"
abbr py python3
abbr rb reboot
abbr rem "sudo apt purge"
abbr ress "sudo systemctl daemon-reload"
abbr rss "sudo systemctl restart"
abbr sch "sudo apt search"
abbr sd "shutdown now"
abbr sipip "sudo pip install"
abbr sss "sudo systemctl start"
abbr stas "systemctl status"
abbr sts "sudo systemctl stop"
abbr tf terraform
abbr tg terragrunt
abbr upd "sudo apt update"
abbr upg "sudo apt upgrade"
abbr vmload "sudo vmware-modconfig --console --install-all"
