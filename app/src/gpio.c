#include "gpio.h"

#define GPIO_BASE 0x60000300

#define GPIO_OUT_W1TS \
    (*(volatile unsigned int *)(GPIO_BASE + 0x04))

#define GPIO_OUT_W1TC \
    (*(volatile unsigned int *)(GPIO_BASE + 0x08))

#define GPIO_ENABLE_W1TS \
    (*(volatile unsigned int *)(GPIO_BASE + 0x24))

#define GPIO2 (1U << 2)

void gpio_init(void)
{
    GPIO_ENABLE_W1TS = GPIO2;
    GPIO_OUT_W1TC = GPIO2;
}

void gpio_set(void)
{
    GPIO_OUT_W1TS = GPIO2;
}

void gpio_clear(void)
{
    GPIO_OUT_W1TC = GPIO2;
}

void gpio_toggle(void)
{
    static int state = 0;

    if (state == 0)
    {
        gpio_set();
        state = 1;
    }
    else
    {
        gpio_clear();
        state = 0;
    }
}
