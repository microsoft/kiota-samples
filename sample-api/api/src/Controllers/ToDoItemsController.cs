// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Identity.Web.Resource;
using ToDoApi.Models;

namespace ToDoApi.Controllers
{
    [Authorize]
    [RequiredScope("ToDoItem.ReadWrite")]
    [ApiController]
    [Route("[controller]")]
    public class ToDoItemsController : ControllerBase
    {
        private const string GetToDoItemRouteName = "GetToDoItem";
        private readonly ToDoContext _context;
        private readonly ILogger<ToDoItemsController> _logger;

        public ToDoItemsController(ToDoContext context, ILogger<ToDoItemsController> logger)
        {
            _context = context;
            _logger = logger;
        }

        // GET: /ToDoItems?isComplete=false
        [HttpGet]
        [ProducesResponseType<IEnumerable<ToDoItem>>(StatusCodes.Status200OK, "application/json")]
        public async Task<ActionResult<IEnumerable<ToDoItem>>> GetToDoItemsAsync([FromQuery] bool? isComplete)
        {
            await LogRequestAsync(Request);
            IQueryable<ToDoItem> items = _context.ToDoItems;
            if (isComplete.HasValue)
            {
                items = items.Where(i => i.IsComplete == isComplete.Value);
            }

            return await items.ToListAsync();
        }

        // GET: /ToDoItems/{id}
        [HttpGet("{id}", Name = GetToDoItemRouteName)]
        [ProducesResponseType<ToDoItem>(StatusCodes.Status200OK, "application/json")]
        [ProducesResponseType<ProblemDetails>(StatusCodes.Status404NotFound, "application/problem+json")]
        public async Task<ActionResult<ToDoItem>> GetToDoItemAsync(Guid id)
        {
            await LogRequestAsync(Request);
            var todoItem = await _context.ToDoItems.FindAsync(id);

            if (todoItem == null)
            {
                return NotFound();
            }

            return todoItem;
        }

        // PATCH: /ToDoItems/{id}
        [HttpPatch("{id}")]
        [Consumes("application/json")]
        [ProducesResponseType(StatusCodes.Status204NoContent)]
        [ProducesResponseType<ProblemDetails>(StatusCodes.Status404NotFound, "application/problem+json")]
        public async Task<IActionResult> PatchToDoItemAsync(Guid id, [FromBody] ToDoItemUpdate update)
        {
            await LogRequestAsync(Request);
            var item = await _context.ToDoItems.FindAsync(id);

            if (item == null)
            {
                return NotFound();
            }

            update.ApplyTo(item);

            try
            {
                await _context.SaveChangesAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (!ToDoItemExists(id))
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return NoContent();
        }

        // POST: /ToDoItems
        [HttpPost]
        [Consumes("application/json")]
        [ProducesResponseType<ToDoItem>(StatusCodes.Status201Created, "application/json")]
        public async Task<ActionResult<ToDoItem>> PostToDoItemAsync([FromBody] ToDoItem todoItem)
        {
            await LogRequestAsync(Request);
            _context.ToDoItems.Add(todoItem);
            await _context.SaveChangesAsync();

            return CreatedAtRoute(GetToDoItemRouteName, new { id = todoItem.Id }, todoItem);
        }

        // DELETE: /ToDoItems/{id}
        [HttpDelete("{id}")]
        [ProducesResponseType(StatusCodes.Status204NoContent)]
        [ProducesResponseType<ProblemDetails>(StatusCodes.Status404NotFound, "application/problem+json")]
        public async Task<IActionResult> DeleteToDoItemAsync(Guid id)
        {
            await LogRequestAsync(Request);
            var todoItem = await _context.ToDoItems.FindAsync(id);
            if (todoItem == null)
            {
                return NotFound();
            }

            _context.ToDoItems.Remove(todoItem);
            await _context.SaveChangesAsync();

            return NoContent();
        }

        private bool ToDoItemExists(Guid id)
        {
            return _context.ToDoItems.Any(e => e.Id == id);
        }

        private async Task LogRequestAsync(HttpRequest request)
        {
            var requestInfo = $"Incoming request: \n{request.Method} {request.Path}{request.QueryString}\n";

            if (request.ContentLength > 0 && request.HasJsonContentType())
            {
                request.Body.Position = 0;
                var reader = new StreamReader(request.Body);
                var payload = await reader.ReadToEndAsync();
                requestInfo += $"\n{payload}";
            }

            _logger.LogInformation(requestInfo);
        }
    }
}
