import asyncio

async def task1(what, when):
    await asyncio.sleep(when)
    print(what)

async def task2(what, when, loop):
    await asyncio.sleep(when)
    print(what)
    loop.stop()

async def end(loop, when):
    await asyncio.sleep(when)
    loop.stop()


loop = asyncio.get_event_loop()

loop.create_task(task1('first hello', 2))
loop.create_task(task2('second hello', 4,loop))


loop.run_forever()
loop.close()