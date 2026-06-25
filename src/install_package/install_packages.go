package install_packages

import (
	"os"
	"log"
	"os/exec"
)

func InstallPackages(packages string) {
	// ИСПРАВЛЕНО: Проверяем UID (User ID) пользователя, а не GID (Group ID)
	if os.Getuid() != 0 {
		log.Fatalf("Ошибка: запустите скрипт через sudo (ваша учетная запись не root)")
	}

	// ИСПРАВЛЕНО: Запускаем через bash -c, чтобы строка с пробелами обработалась корректно.
	// Также добавляем флаг --noconfirm, чтобы pacman не ждал ответа [Y/n] в фоне.
	cmd := exec.Command("bash", "-c", "pacman -S --noconfirm "+packages)

	// Направляем стандартный вывод и ошибки команды прямо в вашу консоль,
	// чтобы вы видели процесс установки pacman в реальном времени.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {

		log.Fatalf("Ошибка при установке пакетов: %v", err)
	}
}
