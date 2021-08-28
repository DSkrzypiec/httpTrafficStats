using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace HttpServer.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class TestController : Controller
    {
        [HttpGet]
        public async Task<int> Get()
        {
            await Task.Delay(1000);
            var rnd = new Random();
            return rnd.Next(1000);
        }
    }
}
