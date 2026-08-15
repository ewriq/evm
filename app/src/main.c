#include "gpio.h"

void app_main(void)
{
    gpio_init();

    while (1)
    {
        gpio_toggle();
    }
}
